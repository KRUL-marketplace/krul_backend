package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"log"
	"user-service/client/db"
	"user-service/internal/repository/model"
)

func (r *repo) Create(ctx context.Context, info *model.UserInfo) (string, error) {
	log.Printf("repo")
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(phoneNumberColumn, nameColumn).
		Values(info.PhoneNumber, info.Name).
		Suffix("RETURNING id")
	log.Printf("repo 2")

	query, args, err := builder.ToSql()
	log.Printf("repo 3")
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}
	log.Printf("repo 4")

	if r.db == nil {
		log.Println("db nil")
	}

	if r.db.DB() == nil {
		log.Println("db db nil")
	}

	var id string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return "", err
	}
	log.Printf("repo 5")

	return id, nil
}
