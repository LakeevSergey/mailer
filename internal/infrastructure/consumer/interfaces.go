package consumer

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
)

type Listner interface {
	Listen(ctx context.Context, worker func(context.Context, entity.SendMail) error) error
}
