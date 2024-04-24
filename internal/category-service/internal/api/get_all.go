package api

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/converter"
	desc "github.com/KRUL-marketplace/krul_backend/internal/category-service/pkg/category-service"
	"log"
)

func (i *Implementation) GetAll(ctx context.Context, req *desc.GetAllRequest) (*desc.GetAllResponse, error) {
	categories, err := i.categoryService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var categoriesMessages []*desc.Category

	for _, category := range categories {
		log.Printf("id: %d, name: %s, slug: %s, created_at: %v, updated_at:%v\n",
			category.ID, category.Info.Name, category.Info.Slug, category.CreatedAt, category.UpdatedAt)

		categoryMessage := converter.ToCategoryFromService(category)
		categoriesMessages = append(categoriesMessages, categoryMessage)
	}

	response := &desc.GetAllResponse{
		Categories: categoriesMessages,
	}

	return response, nil
}
