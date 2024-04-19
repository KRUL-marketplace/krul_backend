package converter

import (
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"
	desc "github.com/KRUL-marketplace/krul_backend/internal/product_service/pkg/product-service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProductInfoFromDesc(info *desc.ProductInfo) *model.ProductInfo {
	return &model.ProductInfo{
		Name:        info.Name,
		Slug:        info.Slug,
		Description: info.Description,
		Price:       info.Price,
	}
}

func ToProductFromService(product *model.Product) *desc.Product {
	var updatedAt *timestamppb.Timestamp
	if product.UpdatedAt.Valid {
		updatedAt = timestamppb.New(product.UpdatedAt.Time)
	}

	return &desc.Product{
		Id:        product.ID,
		Info:      ToProductInfoFromService(product.Info),
		CreatedAt: timestamppb.New(product.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToProductInfoFromService(info model.ProductInfo) *desc.ProductInfo {
	return &desc.ProductInfo{
		Name:        info.Name,
		Slug:        info.Slug,
		Description: info.Description,
		Price:       info.Price,
	}
}

func ToProductFromRepo(product *model.Product) *model.Product {
	return &model.Product{
		ID:        product.ID,
		Info:      ToProductInfoFromRepo(product.Info),
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func ToProductInfoFromRepo(info model.ProductInfo) model.ProductInfo {
	return model.ProductInfo{
		Name:        info.Name,
		Slug:        info.Slug,
		Description: info.Description,
		Price:       info.Price,
	}
}
