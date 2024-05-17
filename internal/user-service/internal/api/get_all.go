package api

import (
	"context"
	"log"
	"user-service/internal/converter"
	desc "user-service/pkg/user-service"
)

func (i *Implementation) GetAll(ctx context.Context, _ *desc.GetAllRequest) (*desc.GetAllResponse, error) {
	users, err := i.userService.GetAll(ctx)

	var userMessages []*desc.User

	for _, user := range users {
		log.Printf("id: %s, name: %s, phone_number: %s, created_at: %v, updated_at: %v\n",
			user.ID, user.Info.Name, user.Info.PhoneNumber, user.CreatedAt, user.UpdatedAt)

		userMessage := converter.ToUserFromService(user)
		userMessages = append(userMessages, userMessage)
	}

	response := &desc.GetAllResponse{
		Users: userMessages,
	}

	return response, err
}
