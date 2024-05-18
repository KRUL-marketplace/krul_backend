package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"user-service/client/db"
	"user-service/internal/converter"
	"user-service/internal/repository/model"
)

func (r *repo) GetAll(ctx context.Context) ([]*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, phoneNumberColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetAll",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User

		err = rows.Scan(&user.ID, &user.Info.Name, &user.Info.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, converter.ToUserFromRepo(&user))
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
