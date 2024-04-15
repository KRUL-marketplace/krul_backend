package product

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"
)

func (s *serv) Get(ctx context.Context, id string) (*model.Product, error) {
	product, err := s.productRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
