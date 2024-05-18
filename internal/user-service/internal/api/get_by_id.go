package api

import (
	"context"
	"log"
	"user-service/internal/converter"
	desc "user-service/pkg/user-service"
)

func (i *Implementation) GetById(ctx context.Context, req *desc.GetByIdRequest) (*desc.GetByIdResponse, error) {
	user, err := i.userService.GetById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %s, name: %s, phone_number: %s, created_at: %v, updated_at: %v",
		user.ID, user.Info.Name, user.Info.PhoneNumber, user.CreatedAt, user.UpdatedAt)

	return &desc.GetByIdResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
