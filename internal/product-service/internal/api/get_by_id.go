package api

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/converter"
	desc "github.com/KRUL-marketplace/krul_backend/internal/product_service/pkg/product-service"
	"log"
)

func (i *Implementation) GetById(ctx context.Context, req *desc.GetByIdRequest) (*desc.GetByIdResponse, error) {
	productObj, err := i.productService.GetById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %s, name: %s, slug: %s, description: %s, price: %d, created_at: %v, updated_at:%v\n",
		productObj.ID, productObj.Info.Name, productObj.Info.Slug, productObj.Info.Description, productObj.Info.Description)

	return &desc.GetByIdResponse{
		Product: converter.ToProductFromService(productObj),
	}, nil
}
