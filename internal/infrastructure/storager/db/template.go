package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/storager"
	"github.com/LakeevSergey/mailer/internal/domain/templatemanager/dto"
)

type DBTemplateStorager struct {
	db *sql.DB
}

func NewDBTemplateStorager(db *sql.DB) *DBTemplateStorager {
	return &DBTemplateStorager{db: db}
}

func (t *DBTemplateStorager) GetByCode(ctx context.Context, code string) (entity.Template, error) {
	var template entity.Template

	err := t.db.QueryRowContext(
		ctx,
		"SELECT id, active, code, name, body, title FROM templates WHERE code = ?",
		code,
	).Scan(&template.Id, &template.Active, &template.Code, &template.Name, &template.Body, &template.Title)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Template{}, storager.ErrorEntityNotFound
	} else if err != nil {
		return entity.Template{}, err
	}

	return template, nil
}

func (t *DBTemplateStorager) Search(ctx context.Context, dto dto.Search) (templates []entity.Template, total int, err error) {
	var count int
	err = t.db.QueryRowContext(
		ctx,
		"SELECT COUNT(*) FROM templates",
	).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	rows, err := t.db.QueryContext(
		ctx,
		"SELECT id, active, code, name, body, title FROM templates ORDER BY id DESC LIMIT ? OFFSET ?",
		dto.Limit, dto.Offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	result := make([]entity.Template, 0, dto.Limit)

	for rows.Next() {
		var template entity.Template

		err = rows.Scan(&template.Id, &template.Active, &template.Code, &template.Name, &template.Body, &template.Title)
		if err != nil {
			return nil, 0, err
		}

		result = append(result, template)
	}

	err = rows.Err()
	if err != nil {
		return nil, 0, err
	}

	return result, count, nil
}

func (t *DBTemplateStorager) Get(ctx context.Context, id int64) (entity.Template, error) {
	var template entity.Template

	err := t.db.QueryRowContext(
		ctx,
		"SELECT id, active, code, name, body, title FROM templates WHERE id = ?",
		id,
	).Scan(&template.Id, &template.Active, &template.Code, &template.Name, &template.Body, &template.Title)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Template{}, storager.ErrorEntityNotFound
	} else if err != nil {
		return entity.Template{}, err
	}

	return template, nil
}

func (t *DBTemplateStorager) Add(ctx context.Context, dto dto.Add) (entity.Template, error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return entity.Template{}, err
	}

	var id int64

	err = tx.QueryRowContext(
		ctx,
		"SELECT id FROM templates WHERE code = ?",
		dto.Code,
	).Scan(&id)

	if err == nil {
		tx.Rollback()
		return entity.Template{}, storager.ErrorDuplicate
	} else if !errors.Is(err, sql.ErrNoRows) {
		tx.Rollback()
		return entity.Template{}, err
	}

	res, err := tx.ExecContext(
		ctx, `
		INSERT INTO templates
			(active, code, name, body, title)
		VALUES
			(?, ?, ?, ?, ?)
	`, dto.Active, dto.Code, dto.Name, dto.Body, dto.Title)

	if err != nil {
		tx.Rollback()
		return entity.Template{}, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return entity.Template{}, err
	}

	template := entity.Template{
		Id:     id,
		Active: dto.Active,
		Code:   dto.Code,
		Name:   dto.Name,
		Body:   dto.Body,
		Title:  dto.Title,
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return entity.Template{}, err
	}

	return template, nil
}

func (t *DBTemplateStorager) Update(ctx context.Context, id int64, dto dto.Update) (entity.Template, error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return entity.Template{}, err
	}

	var sameCodeId int64

	err = tx.QueryRowContext(
		ctx,
		"SELECT id FROM templates WHERE code = ?",
		dto.Code,
	).Scan(&sameCodeId)

	if err == nil && sameCodeId != id {
		tx.Rollback()
		return entity.Template{}, storager.ErrorDuplicate
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		tx.Rollback()
		return entity.Template{}, err
	}

	res, err := tx.ExecContext(
		ctx, `
		UPDATE templates SET
			active = ?,
			code = ?,
			name = ?,
			body = ?,
			title = ?
		WHERE
			id = ?
	`, dto.Active, dto.Code, dto.Name, dto.Body, dto.Title, id)
	if err != nil {
		tx.Rollback()
		return entity.Template{}, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return entity.Template{}, err
	}

	if count == 0 {
		tx.Rollback()
		return entity.Template{}, storager.ErrorEntityNotFound
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return entity.Template{}, err
	}

	template := entity.Template{
		Id:     id,
		Active: dto.Active,
		Code:   dto.Code,
		Name:   dto.Name,
		Body:   dto.Body,
		Title:  dto.Title,
	}

	return template, nil
}

func (t *DBTemplateStorager) Delete(ctx context.Context, id int64) error {
	res, err := t.db.ExecContext(ctx, "DELETE FROM templates WHERE id = ?", id)

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
