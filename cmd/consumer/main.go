package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/LakeevSergey/mailer/internal/application/attachmentmanager"
	"github.com/LakeevSergey/mailer/internal/application/requestprocessor"
	"github.com/LakeevSergey/mailer/internal/config"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/infrastructure/consumer"
	"github.com/LakeevSergey/mailer/internal/infrastructure/consumer/builder"
	"github.com/LakeevSergey/mailer/internal/infrastructure/consumer/sender"
	"github.com/LakeevSergey/mailer/internal/infrastructure/filestorager"
	"github.com/LakeevSergey/mailer/internal/infrastructure/logger"
	"github.com/LakeevSergey/mailer/internal/infrastructure/queue"
	"github.com/LakeevSergey/mailer/internal/infrastructure/queue/encoder"
	"github.com/LakeevSergey/mailer/internal/infrastructure/storager/db"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

	fileinfoStorager := db.NewDBFileInfoStorager(dbMysql)
	filestorager := filestorager.NewLocalFileStorager("./uploads/", func() string { return uuid.NewString() })
	attachmentManager := attachmentmanager.NewAttachmentManager(fileinfoStorager, filestorager)

	requestprocessor := requestprocessor.NewSendMailRequestProcessor(
		templateStorager,
		builder,
		sender,
		attachmentManager,
		entity.SendFrom{Name: cfg.SendFromName, Email: cfg.SendFromEmail},
	)

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
		RetryCount:    cfg.RetryCount,
	}
	queue, err := queue.NewRabbitMQ[entity.SendMail](rbmqConfig, encoder, logger)
	if err != nil {
		logger.ErrorErr(fmt.Errorf("declare RBMQ queue error: %w", err))
		return
	}
	defer queue.Close()

	consumer := consumer.NewConsumer(requestprocessor, queue, logger)

	err = consumer.Run(ctx)
	if err != nil {
		logger.ErrorErr(fmt.Errorf("run consumer error: %w", err))
		return
	}
	<-ctx.Done()
}
