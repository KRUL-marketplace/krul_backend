package service

import (
	"context"
	"user-service/internal/repository/model"
)

func (s *userService) GetAll(ctx context.Context) ([]*model.User, error) {
	users, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
