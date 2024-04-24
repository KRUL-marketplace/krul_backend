package service

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/repository/category/model"
)

type CategoryService interface {
	Create(ctx context.Context, info *model.CategoryInfo) (uint32, error)
	Update(ctx context.Context, id uint32, info *model.CategoryInfo) error
	GetAll(ctx context.Context) ([]*model.Category, error)
	GetById(ctx context.Context, id uint32) (*model.Category, error)
	GetBySlug(ctx context.Context, slug string) (*model.Category, error)
	DeleteById(ctx context.Context, id uint32) error
}
