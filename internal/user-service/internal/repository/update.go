package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"time"
	"user-service/client/db"
	"user-service/internal/repository/model"
)

func (r *repo) Update(ctx context.Context, id string, info *model.UserInfo) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, info.Name).
		Set(phoneNumberColumn, info.PhoneNumber).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
