package api

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/converter"
	desc "github.com/KRUL-marketplace/krul_backend/internal/product_service/pkg/product-service"
	"log"
)

func (i *Implementation) GetBySlug(ctx context.Context, req *desc.GetBySlugRequest) (*desc.GetBySlugResponse, error) {
	productObj, err := i.productService.GetBySlug(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %s, name: %s, slug: %s, description: %s, price: %d, created_at: %v, updated_at:%v\n",
		productObj.ID, productObj.Info.Name, productObj.Info.Slug, productObj.Info.Description, productObj.Info.Description)

	return &desc.GetBySlugResponse{
		Product: converter.ToProductFromService(productObj),
	}, nil
}
