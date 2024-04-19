package product

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"
	"github.com/gosimple/slug"
)

func (s *serv) Create(ctx context.Context, info *model.ProductInfo) (string, error) {
	info.Slug = slug.Make(info.Name)

	var id string
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.productRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return id, nil
}
