package api

import (
	categoryService "product-catalog-service/internal/service/category"
	productService "product-catalog-service/internal/service/product"
	desc "product-catalog-service/pkg/product-catalog-service"
)

type Implementation struct {
	desc.UnimplementedProductCatalogServiceServer
	productService  productService.ProductService
	categoryService categoryService.CategoryService
}

func NewImplementation(productService productService.ProductService, categoryService categoryService.CategoryService) *Implementation {
	return &Implementation{
		productService:  productService,
		categoryService: categoryService,
	}
}
