package api

import (
	"context"
	desc "github.com/KRUL-marketplace/krul_backend/internal/category-service/pkg/category-service"
	"log"
)

func (i *Implementation) DeleteById(ctx context.Context, req *desc.DeleteByIdRequest) (*desc.DeleteByIdResponse, error) {
	err := i.categoryService.DeleteById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("product deleted by id %d", req.GetId())

	return &desc.DeleteByIdResponse{
		Message: "Success",
	}, nil
}
