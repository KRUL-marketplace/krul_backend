package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"user-service/internal/repository/model"
	desc "user-service/pkg/user-service"
)

func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:        info.Name,
		PhoneNumber: info.PhoneNumber,
	}
}

func ToUserFromRepo(user *model.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserInfoFromRepo(info model.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:        info.Name,
		PhoneNumber: info.PhoneNumber,
	}
}

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromService(info model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Name:        info.Name,
		PhoneNumber: info.PhoneNumber,
	}
}
