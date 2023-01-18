package consumer

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/requestprocessor"
	"github.com/LakeevSergey/mailer/internal/infrastructure"
)

type Consumer struct {
	requestProcessor requestprocessor.SendMailRequestProcessor
	listner          Listner
	logger           infrastructure.Logger
}

func NewConsumer(requestProcessor requestprocessor.SendMailRequestProcessor, listner Listner, logger infrastructure.Logger) *Consumer {
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
