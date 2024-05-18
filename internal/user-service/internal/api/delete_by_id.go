package api

import (
	"context"
	"log"
	desc "user-service/pkg/user-service"
)

func (i *Implementation) DeleteById(ctx context.Context, req *desc.DeleteByIdRequest) (*desc.DeleteByIdResponse, error) {
	err := i.userService.DeleteById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("user deleted by id")

	return &desc.DeleteByIdResponse{
		Message: "Success",
	}, nil
}
