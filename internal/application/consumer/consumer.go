package consumer

import "context"

type Consumer struct {
	requestProcessor SendMailRequestProcessor
	listner          Listner
}

func NewConsumer(requestProcessor SendMailRequestProcessor, listner Listner) *Consumer {
	return &Consumer{
		requestProcessor: requestProcessor,
		listner:          listner,
	}
}

func (c *Consumer) Run(ctx context.Context) error {
	return c.listner.Listen(ctx, c.requestProcessor.Process)
}
