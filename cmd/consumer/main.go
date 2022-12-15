package main

import (
	"context"

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

	templateStorager := db.NewDBTemplateStorager()
	builder := builder.NewTwigBuilder()
	sender := sender.NewSMTPSender()

	mailer := domain.NewMailer(templateStorager, builder, sender, entity.SendFrom{Name: "test", Email: "test@test.test"})

	decoder := encoder.NewJSONEncoder[dto.SendMail]()
	listner := listner.NewRabbitMQListner[dto.SendMail](decoder)

	consumer := consumer.NewConsumer(mailer, listner)

	consumer.Run(ctx)
}
