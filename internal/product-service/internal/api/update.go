package api

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/converter"
	desc "github.com/KRUL-marketplace/krul_backend/internal/product_service/pkg/product-service"
	"log"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.CreateResponse, error) {
	id := req.GetId()
	resId, err := i.productService.Update(ctx, id, converter.ToProductInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("updated product with id: %s", resId)

	return &desc.CreateResponse{
		Id: resId,
	}, nil
}
