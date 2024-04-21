package api

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/converter"
	desc "github.com/KRUL-marketplace/krul_backend/internal/category-service/pkg/category-service"
	"log"
)

func (i *Implementation) GetById(ctx context.Context, req *desc.GetByIdRequest) (*desc.GetByIdResponse, error) {
	category, err := i.categoryService.GetById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, slug: %s, created_at: %v, updated_at: %v",
		category.ID, category.Info.Name, category.Info.Slug, category.CreatedAt, category.UpdatedAt)

	return &desc.GetByIdResponse{
		Category: converter.ToCategoryFromService(category),
	}, nil
}
