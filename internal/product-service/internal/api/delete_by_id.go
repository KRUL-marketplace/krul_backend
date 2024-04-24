package api

import (
	"context"
	desc "github.com/KRUL-marketplace/krul_backend/internal/product_service/pkg/product-service"
	"log"
)

func (i *Implementation) DeleteById(ctx context.Context, req *desc.DeleteByIdRequest) (*desc.DeleteByIdResponse, error) {
	err := i.productService.DeleteById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("product deleted by id")

	return &desc.DeleteByIdResponse{
		Message: "Success",
	}, nil
}
