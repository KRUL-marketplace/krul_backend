package product

import (
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/service"
)

type serv struct {
	productRepository repository.ProductRepository
}

func NewService(productRepository repository.ProductRepository) service.ProductService {
	return &serv{}
}
