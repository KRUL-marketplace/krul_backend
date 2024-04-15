package product

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"
)

func (s *serv) GetAll(ctx context.Context) ([]*model.Product, error) {
	products, err := s.productRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}
