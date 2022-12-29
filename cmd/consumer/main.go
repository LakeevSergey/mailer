package main

import (
	"context"
	"fmt"
	"os"

	"github.com/LakeevSergey/mailer/internal/application/config"
	"github.com/LakeevSergey/mailer/internal/application/consumer"
	"github.com/LakeevSergey/mailer/internal/application/dto"
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
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		fmt.Printf("Parse config error: %v", err)
		return
	}
	zlogger := zerolog.New(os.Stdout).Level(zerolog.Level(cfg.ConsoleLoggerLevel))
	logger := logger.NewLogger(zlogger)

	templateStorager := db.NewDBTemplateStorager()
	builder := builder.NewTwigBuilder()
	sender := sender.NewSMTPSender()
	mailer := mailer.NewMailer(templateStorager, builder, sender, entity.SendFrom{Name: cfg.SendFromName, Email: cfg.SendFromEmail})

	coder := coder.NewJSONCoder[dto.SendMail]()
	queue := queue.NewRabbitMQ[dto.SendMail](coder, logger)

	consumer := consumer.NewConsumer(mailer, queue)

	consumer.Run(ctx)
}
