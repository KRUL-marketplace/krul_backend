package converter

import (
	"cart-service/internal/repository/model"
	desc "cart-service/pkg/cart-service"
)

func ToCartProductInfoFromDesc(info *desc.CartProductInfo) *model.CartProductInfo {
	cartProductInfo := &model.CartProductInfo{
		ProductId: info.ProductID,
		Name:      info.Name,
		Slug:      info.Slug,
		Image:     info.Image,
		Price:     info.Price,
		Quantity:  info.Quantity,
	}

	return cartProductInfo
}
