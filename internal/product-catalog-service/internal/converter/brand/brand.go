package brand

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"product-catalog-service/internal/repository/brand/model"
	desc "product-catalog-service/pkg/product-catalog-service"
)

func ToBrandInfoFromDesc(info *desc.BrandInfo) *model.BrandInfo {
	return &model.BrandInfo{
		Name: info.Name,
		Slug: info.Slug,
	}
}

func ToBrandFromService(category *model.Brand) *desc.Brand {
	var updatedAt *timestamppb.Timestamp
	if category.UpdatedAt.Valid {
		updatedAt = timestamppb.New(category.UpdatedAt.Time)
	}

	return &desc.Brand{
		Id:        category.ID,
		Info:      ToBrandInfoFromService(category.Info),
		CreatedAt: timestamppb.New(category.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToBrandInfoFromService(info model.BrandInfo) *desc.BrandInfo {
	return &desc.BrandInfo{
		Name: info.Name,
		Slug: info.Slug,
	}
}

func ToBrandFromRepo(category *model.Brand) *model.Brand {
	return &model.Brand{
		ID:        category.ID,
		Info:      ToBrandInfoFromRepo(category.Info),
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func ToBrandInfoFromRepo(info model.BrandInfo) model.BrandInfo {
	return model.BrandInfo{
		Name: info.Name,
		Slug: info.Slug,
	}
}
