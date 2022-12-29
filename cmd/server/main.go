package main

import (
	"context"
	"fmt"
	"os"

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
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		fmt.Printf("Parse config error: %v", err)
		return
	}
	zlogger := zerolog.New(os.Stdout).Level(zerolog.Level(cfg.ConsoleLoggerLevel))
	logger := logger.NewLogger(zlogger)

	coder := coder.NewJSONCoder[entity.SendMail]()
	queue := queue.NewRabbitMQ[entity.SendMail](coder, logger)
	mailSender := mailsender.NewMailSender(queue)
	api := api.NewJSONApi(coder, mailSender)
	router := router.NewRouter(api, logger)
	server := server.NewServer(fmt.Sprintf(":%d", cfg.ApiPort), router)

	server.Run(ctx)
}
