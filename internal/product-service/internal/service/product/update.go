package product

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"
	"github.com/gosimple/slug"
)

func (s *serv) Update(ctx context.Context, id string, info *model.ProductInfo) (string, error) {
	info.Slug = slug.Make(info.Name)

	var resId string
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		resId, errTx = s.productRepository.Update(ctx, id, info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return resId, nil
}
