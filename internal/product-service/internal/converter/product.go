package converter

import (
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"
	desc "github.com/KRUL-marketplace/krul_backend/internal/product_service/pkg/product-service"
)

func ToProductInfoFromDesc(info *desc.ProductInfo) *model.ProductInfo {
	return &model.ProductInfo{
		Name:        info.Name,
		Description: info.Description,
		Price:       info.Price,
	}
}
