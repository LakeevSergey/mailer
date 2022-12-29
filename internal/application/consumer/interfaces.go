package consumer

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
)

type SendMailRequestProcessor interface {
	Process(ctx context.Context, sendMail entity.SendMail) error
}

type Listner interface {
	Listen(ctx context.Context, worker func(context.Context, entity.SendMail) error) error
}
