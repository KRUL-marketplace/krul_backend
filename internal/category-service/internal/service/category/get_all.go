package category

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/repository/category/model"
)

func (s *serv) GetAll(ctx context.Context) ([]*model.Category, error) {
	categories, err := s.categoryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
