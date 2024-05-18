package service

import (
	"context"
	"user-service/client/db"
	"user-service/internal/repository"
	"user-service/internal/repository/model"
)

type userService struct {
	userRepository repository.Repository
	txManager      db.TxManager
}

type UserService interface {
	Create(ctx context.Context, user *model.UserInfo) (string, error)
	Update(ctx context.Context, id string, user *model.UserInfo) error
	GetAll(ctx context.Context) ([]*model.User, error)
	GetById(ctx context.Context, id string) (*model.User, error)
	DeleteById(ctx context.Context, id string) error
}

func NewService(userRepository repository.Repository, txManager db.TxManager) UserService {
	return &userService{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
