package api

import (
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/service"
	desc "github.com/KRUL-marketplace/krul_backend/internal/product_service/pkg/product-service"
)

type Implementation struct {
	desc.UnimplementedProductServiceServer
	productService service.ProductService
}

func NewImplementation(productService service.ProductService) *Implementation {
	return &Implementation{
		productService: productService,
	}
}
