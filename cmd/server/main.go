package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LakeevSergey/mailer/internal/common/config"
	"github.com/LakeevSergey/mailer/internal/common/dto"
	"github.com/LakeevSergey/mailer/internal/common/encoder"
	"github.com/LakeevSergey/mailer/internal/server"
	"github.com/LakeevSergey/mailer/internal/server/api"
	"github.com/LakeevSergey/mailer/internal/server/domain/mailsender"
	"github.com/LakeevSergey/mailer/internal/server/requestsavier"
	"github.com/LakeevSergey/mailer/internal/server/router"
)

func main() {
	ctx := context.Background()
	logger := log.Default()
	cfg, err := config.New()
	if err != nil {
		logger.Printf("Parse config error: %v", err)
		return
	}

	encoder := encoder.NewJSONEncoder[dto.SendMail]()
	savier := requestsavier.NewRabbitMQRequestSavier[dto.SendMail](encoder)
	mailSender := mailsender.NewMailSender(savier)
	api := api.NewJSONApi(encoder, mailSender)
	router := router.NewRouter(api, logger)
	server := server.NewServer(fmt.Sprintf(":%d", cfg.ApiPort), router)

	server.Run(ctx)
}
