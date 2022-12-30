package consumer

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/application"
)

type Consumer struct {
	requestProcessor SendMailRequestProcessor
	listner          Listner
	logger           application.Logger
}

func NewConsumer(requestProcessor SendMailRequestProcessor, listner Listner, logger application.Logger) *Consumer {
	return &Consumer{
		requestProcessor: requestProcessor,
		listner:          listner,
		logger:           logger,
	}
}

func (c *Consumer) Run(ctx context.Context) {
	go c.logger.ErrorErr(c.listner.Listen(ctx, c.requestProcessor.Process))
	c.logger.Info("Consumer running")
}
