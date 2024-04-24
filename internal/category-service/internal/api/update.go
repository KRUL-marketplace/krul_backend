package api

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/converter"
	desc "github.com/KRUL-marketplace/krul_backend/internal/category-service/pkg/category-service"
	"log"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	id := req.GetId()
	err := i.categoryService.Update(ctx, id, converter.ToCategoryInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("updated product with id: %d", id)

	return &desc.UpdateResponse{
		Message: "Success",
	}, nil
}
