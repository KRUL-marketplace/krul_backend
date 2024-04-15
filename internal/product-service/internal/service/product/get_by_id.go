package product

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"
)

func (s *serv) GetById(ctx context.Context, id string) (*model.Product, error) {
	product, err := s.productRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
