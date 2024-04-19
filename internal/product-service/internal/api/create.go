package api

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/converter"
	desc "github.com/KRUL-marketplace/krul_backend/internal/product_service/pkg/product-service"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.productService.Create(ctx, converter.ToProductInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted product with id: %s", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
