package consumer

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/common/dto"
)

type Mailer interface {
	Send(sendMail dto.SendMail) error
}

type Listner interface {
	Listen(ctx context.Context, worker func(dto.SendMail) error) error
}
