package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/LakeevSergey/mailer/internal/application/config"
	"github.com/LakeevSergey/mailer/internal/application/consumer"
	"github.com/LakeevSergey/mailer/internal/application/logger"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/mailer"
	"github.com/LakeevSergey/mailer/internal/infrastructure/builder"
	"github.com/LakeevSergey/mailer/internal/infrastructure/coder"
	"github.com/LakeevSergey/mailer/internal/infrastructure/queue"
	"github.com/LakeevSergey/mailer/internal/infrastructure/sender"
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
	builder := builder.NewTwigBuilder()
	sender := sender.NewSMTPSender(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPassword)
	mailer := mailer.NewMailer(templateStorager, builder, sender, entity.SendFrom{Name: cfg.SendFromName, Email: cfg.SendFromEmail})

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
		RetryCount:    cfg.RetryCount,
	}
	queue, err := queue.NewRabbitMQ[entity.SendMail](rbmqConfig, coder, logger)
	if err != nil {
		logger.ErrorErr(fmt.Errorf("declare RBMQ queue error: %w", err))
		return
	}
	defer queue.Close()

	consumer := consumer.NewConsumer(mailer, queue, logger)

	err = consumer.Run(ctx)
	if err != nil {
		logger.ErrorErr(fmt.Errorf("run consumer error: %w", err))
		return
	}
	<-ctx.Done()
}
