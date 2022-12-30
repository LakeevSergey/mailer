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

func (c *Consumer) Run(ctx context.Context) error {
	c.logger.Info("Consumer is running")
	return c.listner.Listen(ctx, c.requestProcessor.Process)
}
