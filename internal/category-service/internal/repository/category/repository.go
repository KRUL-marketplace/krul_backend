package category

import (
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/client/db"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.CategoryRepository {
	return &repo{
		db: db,
	}
}
