package product

import (
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/client/db"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/service"
)

type serv struct {
	productRepository repository.ProductRepository
	txManager         db.TxManager
}

func NewService(productRepository repository.ProductRepository, txManager db.TxManager) service.ProductService {
	return &serv{
		productRepository: productRepository,
		txManager:         txManager,
	}
}
