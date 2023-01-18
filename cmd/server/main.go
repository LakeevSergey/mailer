package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/LakeevSergey/mailer/internal/application/attachmentmanager"
	"github.com/LakeevSergey/mailer/internal/application/mailsender"
	"github.com/LakeevSergey/mailer/internal/application/templatemanager"
	"github.com/LakeevSergey/mailer/internal/config"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/infrastructure/filestorager"
	"github.com/LakeevSergey/mailer/internal/infrastructure/logger"
	"github.com/LakeevSergey/mailer/internal/infrastructure/queue"
	"github.com/LakeevSergey/mailer/internal/infrastructure/queue/encoder"
	"github.com/LakeevSergey/mailer/internal/infrastructure/server"
	apijson "github.com/LakeevSergey/mailer/internal/infrastructure/server/api/json"
	"github.com/LakeevSergey/mailer/internal/infrastructure/server/hasher"
	"github.com/LakeevSergey/mailer/internal/infrastructure/server/router"
	"github.com/LakeevSergey/mailer/internal/infrastructure/server/router/middleware"
	"github.com/LakeevSergey/mailer/internal/infrastructure/server/sign"
	"github.com/LakeevSergey/mailer/internal/infrastructure/storager/db"
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

	encoder := encoder.NewJSONEncoder[entity.SendMail]()

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
	queue, err := queue.NewRabbitMQ[entity.SendMail](rbmqConfig, encoder, logger)
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

	fileinfoStorager := db.NewDBFileInfoStorager(dbMysql)
	filestorager := filestorager.NewLocalFileStorager("./uploads/", func() string { return uuid.NewString() })
	attachmentmanager := attachmentmanager.NewAttachmentManager(fileinfoStorager, filestorager)

	api := apijson.NewJSONApi(mailSender, templateManager, attachmentmanager)

	var hashChecker sign.HashChecker
	if cfg.HasherKey != "" {
		hashChecker = hasher.NewSha256Hasher(cfg.HasherKey)
	} else {
		hashChecker = hasher.NewEmptyHasher()
	}

	handlers := [](func(http.Handler) http.Handler){}
	headersToCheck := []string{}
	if cfg.TimestampHeader != "" {
		timestampChecker := middleware.NewTimestampChecker(cfg.TimestampHeader, cfg.TimestampDelay)
		handlers = append(handlers, timestampChecker.Handler)
		headersToCheck = append(headersToCheck, cfg.TimestampHeader)
	}

	if cfg.SignatureHeader != "" {
		signChecker := sign.NewHTTPRequestSignChecker(hashChecker, cfg.SignatureHeader, headersToCheck)
		handlers = append(handlers, middleware.NewSignCheckerMiddleware(signChecker).Handler)
	}

	router := router.NewRouter(api, logger, handlers...)
	server := server.NewServer(fmt.Sprintf(":%d", cfg.ApiPort), router, logger)

	err = server.Run(ctx)
	if err != nil {
		logger.ErrorErr(fmt.Errorf("run server error: %w", err))
		return
	}
	<-ctx.Done()
}
