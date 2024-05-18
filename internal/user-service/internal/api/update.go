package api

import (
	"context"
	"log"
	"user-service/internal/converter"
	desc "user-service/pkg/user-service"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	id := req.GetId()
	err := i.userService.Update(ctx, id, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("updated user with id %s", id)

	return &desc.UpdateResponse{
		Id: id,
	}, nil
}
