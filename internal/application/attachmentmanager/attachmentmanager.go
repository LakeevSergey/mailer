package attachmentmanager

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/attachmentmanager/dto"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/storager"
	storagerdto "github.com/LakeevSergey/mailer/internal/domain/storager/dto"
)

type AttachmentManager struct {
	fileInfoStorager storager.FileinfoStorager
	fileStorager     FileStorager
}

func NewAttachmentManager(fileInfoStorager storager.FileinfoStorager, fileStorager FileStorager) *AttachmentManager {
	return &AttachmentManager{
		fileInfoStorager: fileInfoStorager,
		fileStorager:     fileStorager,
	}
}

func (s *AttachmentManager) Add(ctx context.Context, dto dto.Add) (int64, int, error) {
	filename, size, err := s.fileStorager.Save(dto.Content)
	if err != nil {
		return 0, 0, err
	}

	fileinfo, err := s.fileInfoStorager.Add(ctx, storagerdto.AddFileInfo{
		FileName: dto.FileName,
		Mime:     dto.Mime,
		Path:     filename,
	})
	if err != nil {
		return 0, 0, err
	}

	return fileinfo.Id, size, nil
}

func (s *AttachmentManager) Get(ctx context.Context, id int64) (entity.File, error) {
	fileInfo, err := s.fileInfoStorager.Get(ctx, id)
	if err != nil {
		return entity.File{}, err
	}

	data, err := s.fileStorager.Get(fileInfo.Path)
	if err != nil {
		return entity.File{}, err
	}

	return entity.File{
		Info: fileInfo,
		Data: data,
	}, nil
}

func (s *AttachmentManager) Delete(ctx context.Context, id int64) error {
	fileInfo, err := s.fileInfoStorager.Get(ctx, id)
	if err != nil {
		return err
	}

	err = s.fileStorager.Delete(fileInfo.Path)
	if err != nil {
		return err
	}

	return s.fileInfoStorager.Delete(ctx, id)
}
