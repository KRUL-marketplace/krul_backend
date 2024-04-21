package api

import (
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/service"
	desc "github.com/KRUL-marketplace/krul_backend/internal/category-service/pkg/category-service"
)

type Implementation struct {
	desc.UnimplementedCategoryServiceServer
	categoryService service.CategoryService
}

func NewImplementation(categoryService service.CategoryService) *Implementation {
	return &Implementation{
		categoryService: categoryService,
	}
}
