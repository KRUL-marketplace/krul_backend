package repository

import (
	"context"
	"user-service/client/db"
	"user-service/internal/repository/model"
)

const (
	tableName         = "users"
	idColumn          = "id"
	phoneNumberColumn = "phone_number"
	nameColumn        = "name"
	createdAtColumn   = "created_at"
	updatedAtColumn   = "updated_at"
)

type Repository interface {
	Create(ctx context.Context, info *model.UserInfo) (string, error)
	Update(ctx context.Context, id string, info *model.UserInfo) error
	GetAll(ctx context.Context) ([]*model.User, error)
	GetById(ctx context.Context, id string) (*model.User, error)
	DeleteById(ctx context.Context, id string) error
}

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) Repository {
	return &repo{
		db: db,
	}
}
