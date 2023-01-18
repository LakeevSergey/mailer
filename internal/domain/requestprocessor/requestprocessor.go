package requestprocessor

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
)

type SendMailRequestProcessor interface {
	Process(ctx context.Context, sendMail entity.SendMail) error
}
