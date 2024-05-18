package api

import (
	"user-service/internal/service"
	desc "user-service/pkg/user-service"
)

type Implementation struct {
	desc.UnimplementedUserServiceServer
	userService service.UserService
}

func NewImplementation(
	userService service.UserService,
) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
