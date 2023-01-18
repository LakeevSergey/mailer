package storager

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/storager/dto"
)

type FileinfoStorager interface {
	Get(ctx context.Context, id int64) (entity.FileInfo, error)
	Add(ctx context.Context, dto dto.AddFileInfo) (entity.FileInfo, error)
	Delete(ctx context.Context, id int64) error
}
