package category

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/repository/category/model"
)

func (s *serv) GetBySlug(ctx context.Context, slug string) (*model.Category, error) {
	category, err := s.categoryRepository.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return category, nil
}
