package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/LakeevSergey/mailer/internal/application/api"
	"github.com/LakeevSergey/mailer/internal/application/config"
	"github.com/LakeevSergey/mailer/internal/application/logger"
	"github.com/LakeevSergey/mailer/internal/application/router"
	"github.com/LakeevSergey/mailer/internal/application/server"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/mailsender"
	"github.com/LakeevSergey/mailer/internal/infrastructure/coder"
	"github.com/LakeevSergey/mailer/internal/infrastructure/queue"
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
	queue := queue.NewRabbitMQ[entity.SendMail](rbmqConfig, coder, logger)

	mailSender := mailsender.NewMailSender(queue)
	api := api.NewJSONApi(coder, mailSender)
	router := router.NewRouter(api, logger)
	server := server.NewServer(fmt.Sprintf(":%d", cfg.ApiPort), router, logger)

	server.Run(ctx)
	<-ctx.Done()
}
