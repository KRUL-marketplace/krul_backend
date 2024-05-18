package api

import (
	"context"
	"log"
	"user-service/internal/converter"
	desc "user-service/pkg/user-service"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateUserInfoRequest) (*desc.CreateUserInfoResponse, error) {
	log.Printf("create user request: %+v", req)
	id, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("created user with id %s", id)

	return &desc.CreateUserInfoResponse{
		Id: id,
	}, nil
}
