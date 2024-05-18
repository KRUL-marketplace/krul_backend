package service

import (
	"context"
	"user-service/internal/repository/model"
)

func (s *userService) GetById(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
