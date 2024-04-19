package product

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"
)

const (
	tableName = "products"
)

type repo struct {
}

func NewRepository() repository.ProductRepository {
	return &repo{}
}

func (r *repo) Create(ctx context.Context, product *model.ProductInfo) (string, error) {
	return "", nil
}
