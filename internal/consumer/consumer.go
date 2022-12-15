package consumer

import "context"

type Consumer struct {
	mailer  Mailer
	listner Listner
}

func NewConsumer(mailer Mailer, listner Listner) *Consumer {
	return &Consumer{
		mailer:  mailer,
		listner: listner,
	}
}

func (c *Consumer) Run(ctx context.Context) error {
	return c.listner.Listen(ctx, c.mailer.Send)
}
