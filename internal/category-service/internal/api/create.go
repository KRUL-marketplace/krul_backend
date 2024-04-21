package api

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/converter"
	desc "github.com/KRUL-marketplace/krul_backend/internal/category-service/pkg/category-service"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.categoryService.Create(ctx, converter.ToCategoryInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted product with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
