package product

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"
)

func (s *serv) GetBySlug(ctx context.Context, slug string) (*model.Product, error) {
	product, err := s.productRepository.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return product, nil
}
