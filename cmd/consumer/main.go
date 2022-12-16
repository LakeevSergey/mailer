package main

import (
	"context"
	"log"

	"github.com/LakeevSergey/mailer/internal/common/config"
	"github.com/LakeevSergey/mailer/internal/common/dto"
	"github.com/LakeevSergey/mailer/internal/common/encoder"
	"github.com/LakeevSergey/mailer/internal/consumer"
	"github.com/LakeevSergey/mailer/internal/consumer/builder"
	"github.com/LakeevSergey/mailer/internal/consumer/domain"
	"github.com/LakeevSergey/mailer/internal/consumer/domain/entity"
	"github.com/LakeevSergey/mailer/internal/consumer/listner"
	"github.com/LakeevSergey/mailer/internal/consumer/sender"
	"github.com/LakeevSergey/mailer/internal/consumer/storager/db"
)

func main() {
	ctx := context.Background()
	logger := log.Default()
	cfg, err := config.New()
	if err != nil {
		logger.Printf("Parse config error: %v", err)
		return
	}

	templateStorager := db.NewDBTemplateStorager()
	builder := builder.NewTwigBuilder()
	sender := sender.NewSMTPSender()

	mailer := domain.NewMailer(templateStorager, builder, sender, entity.SendFrom{Name: cfg.SendFromName, Email: cfg.SendFromEmail})

	decoder := encoder.NewJSONEncoder[dto.SendMail]()
	listner := listner.NewRabbitMQListner[dto.SendMail](decoder)

	consumer := consumer.NewConsumer(mailer, listner)

	consumer.Run(ctx)
}
