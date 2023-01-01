package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	apijson "github.com/LakeevSergey/mailer/internal/application/api/json"
	"github.com/LakeevSergey/mailer/internal/application/config"
	"github.com/LakeevSergey/mailer/internal/application/logger"
	"github.com/LakeevSergey/mailer/internal/application/router"
	"github.com/LakeevSergey/mailer/internal/application/server"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/mailsender"
	"github.com/LakeevSergey/mailer/internal/domain/templatemanager"
	"github.com/LakeevSergey/mailer/internal/infrastructure/coder"
	"github.com/LakeevSergey/mailer/internal/infrastructure/queue"
	"github.com/LakeevSergey/mailer/internal/infrastructure/storager/db"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	cfg, err := config.New()
	if err != nil {
		fmt.Printf("Parse config error: %v", err)
		return
	}
	zlogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.Level(cfg.ConsoleLoggerLevel)).With().Timestamp().Logger()
	logger := logger.NewLogger(zlogger)

	coder := coder.NewJSONCoder[entity.SendMail]()

	rbmqConfig := queue.Config{
		User:          cfg.RBMQUser,
		Password:      cfg.RBMQPassword,
		Host:          cfg.RBMQHost,
		Port:          cfg.RBMQPort,
		Queue:         cfg.RBMQQueue,
		ExchangeInput: cfg.RBMQExchangeInput,
		ExchangeDLX:   cfg.RBMQExchangeDLX,
		QueueDLX:      cfg.RBMQQueueDLX,
		RetryDelay:    cfg.RetryDelay,
	}
	queue, err := queue.NewRabbitMQ[entity.SendMail](rbmqConfig, coder, logger)
	if err != nil {
		logger.ErrorErr(fmt.Errorf("declare RBMQ queue error: %w", err))
		return
	}
	defer queue.Close()

	dbconfig := mysql.NewConfig()
	dbconfig.Net = "tcp"
	dbconfig.Addr = cfg.DBHost + ":" + strconv.Itoa(cfg.DBPort)
	dbconfig.User = cfg.DBUser
	dbconfig.Passwd = cfg.DBPassword
	dbconfig.DBName = cfg.DBName

	dbMysql, err := sql.Open("mysql", dbconfig.FormatDSN())
	if err != nil {
		logger.ErrorErr(fmt.Errorf("open db connection error: %w", err))
		return
	}
	defer dbMysql.Close()

	templateStorager := db.NewDBTemplateStorager(dbMysql)

	mailSender := mailsender.NewMailSender(queue)
	templateManager := templatemanager.NewTemplateManager(templateStorager)
	api := apijson.NewJSONApi(mailSender, templateManager)
	router := router.NewRouter(api, logger)
	server := server.NewServer(fmt.Sprintf(":%d", cfg.ApiPort), router, logger)

	err = server.Run(ctx)
	if err != nil {
		logger.ErrorErr(fmt.Errorf("run server error: %w", err))
		return
	}
	<-ctx.Done()
}
