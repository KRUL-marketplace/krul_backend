package service

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"
)

type ProductService interface {
	Create(ctx context.Context, info *model.ProductInfo) (string, error)
	GetById(ctx context.Context, id string) (*model.Product, error)
	GetAll(ctx context.Context) ([]*model.Product, error)
}
