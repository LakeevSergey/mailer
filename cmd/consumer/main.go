package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
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

	templateStorager := db.NewDBTemplateStorager()
	builder := builder.NewTwigBuilder()
	sender := sender.NewSMTPSender()
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
	}
	queue := queue.NewRabbitMQ[entity.SendMail](rbmqConfig, coder, logger)

	consumer := consumer.NewConsumer(mailer, queue, logger)

	consumer.Run(ctx)
	<-ctx.Done()
}
