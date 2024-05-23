package api

import (
	"cart-service/internal/converter"
	desc "cart-service/pkg/cart-service"
	"context"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.AddProductRequest) (*desc.AddProductResponse, error) {
	id, err := i.cartService.Delete(ctx, req.GetUserId(), converter.ToCartProductInfoFromDesc(req.GetProduct()))
	if err != nil {
		return nil, err
	}

	return &desc.AddProductResponse{
		Id: id,
	}, nil
}
