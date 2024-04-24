package category

import (
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/client/db"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/repository"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/service"
)

type serv struct {
	categoryRepository repository.CategoryRepository
	txManager          db.TxManager
}

func NewService(categoryRepository repository.CategoryRepository, txManager db.TxManager) service.CategoryService {
	return &serv{
		categoryRepository: categoryRepository,
		txManager:          txManager,
	}
}
