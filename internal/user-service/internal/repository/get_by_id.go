package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"user-service/client/db"
	"user-service/internal/converter"
	"user-service/internal/repository/model"
)

func (r *repo) GetById(ctx context.Context, id string) (*model.User, error) {
	builder := sq.Select(idColumn, phoneNumberColumn, nameColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetById",
		QueryRaw: query,
	}

	var user model.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).
		Scan(&user.ID, &user.Info.Name, &user.Info.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
