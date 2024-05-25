package api

import (
	"cart-service/internal/converter"
	desc "cart-service/pkg/cart-service"
	"context"
	"log"
)

func (i *Implementation) GetUserCart(ctx context.Context, req *desc.GetUserCartRequest) (*desc.GetUserCartResponse, error) {
	cart, err := i.cartService.GetUserCart(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	log.Printf("[MY] %+v", cart.Products)

	return &desc.GetUserCartResponse{
		Cart: converter.ToCartFromService(cart),
	}, nil
}
