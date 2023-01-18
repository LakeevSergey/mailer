package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/storager"
	"github.com/LakeevSergey/mailer/internal/domain/storager/dto"
)

type DBFileInfoStorager struct {
	db *sql.DB
}

func NewDBFileInfoStorager(db *sql.DB) *DBFileInfoStorager {
	return &DBFileInfoStorager{db: db}
}

func (s *DBFileInfoStorager) Get(ctx context.Context, id int64) (entity.FileInfo, error) {
	var fileinfo entity.FileInfo

	err := s.db.QueryRowContext(
		ctx,
		"SELECT id, filename, mime, path FROM fileinfo WHERE id = ?",
		id,
	).Scan(&fileinfo.Id, &fileinfo.FileName, &fileinfo.Mime, &fileinfo.Path)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.FileInfo{}, storager.ErrorEntityNotFound
	} else if err != nil {
		return entity.FileInfo{}, err
	}

	return fileinfo, nil
}

func (s *DBFileInfoStorager) Add(ctx context.Context, dto dto.AddFileInfo) (entity.FileInfo, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return entity.FileInfo{}, err
	}

	var id int64

	err = tx.QueryRowContext(
		ctx,
		"SELECT id FROM fileinfo WHERE path = ?",
		dto.Path,
	).Scan(&id)

	if err == nil {
		tx.Rollback()
		return entity.FileInfo{}, storager.ErrorDuplicate
	} else if !errors.Is(err, sql.ErrNoRows) {
		tx.Rollback()
		return entity.FileInfo{}, err
	}

	res, err := tx.ExecContext(
		ctx, `
		INSERT INTO fileinfo
			(filename, mime, path)
		VALUES
			(?, ?, ?)
	`, dto.FileName, dto.Mime, dto.Path)

	if err != nil {
		tx.Rollback()
		return entity.FileInfo{}, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return entity.FileInfo{}, err
	}

	fileinfo := entity.FileInfo{
		Id:       id,
		FileName: dto.FileName,
		Mime:     dto.Mime,
		Path:     dto.Path,
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return entity.FileInfo{}, err
	}

	return fileinfo, nil
}

func (s *DBFileInfoStorager) Delete(ctx context.Context, id int64) error {
	res, err := s.db.ExecContext(ctx, "DELETE FROM fileinfo WHERE id = ?", id)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return storager.ErrorEntityNotFound
	}

	return nil
}
