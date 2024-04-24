package category

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/repository/category/model"
)

func (s *serv) GetById(ctx context.Context, id uint32) (*model.Category, error) {
	category, err := s.categoryRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
