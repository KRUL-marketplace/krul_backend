package api

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/converter"
	desc "github.com/KRUL-marketplace/krul_backend/internal/category-service/pkg/category-service"
	"log"
)

func (i *Implementation) GetBySlug(ctx context.Context, req *desc.GetBySlugRequest) (*desc.GetBySlugResponse, error) {
	category, err := i.categoryService.GetBySlug(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, slug: %s, created_at: %v, updated_at: %v",
		category.ID, category.Info.Name, category.Info.Slug, category.CreatedAt, category.UpdatedAt)

	return &desc.GetBySlugResponse{
		Category: converter.ToCategoryFromService(category),
	}, nil
}
