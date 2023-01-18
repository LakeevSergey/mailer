package attachmentmanager

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/attachmentmanager/dto"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
)

type AttachmentManager interface {
	Add(ctx context.Context, file dto.Add) (int64, int, error)
	Get(ctx context.Context, id int64) (entity.File, error)
	Delete(ctx context.Context, id int64) error
}
