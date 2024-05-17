package service

import (
	"context"
	"log"
	"user-service/internal/repository/model"
)

func (s *userService) Create(ctx context.Context, info *model.UserInfo) (string, error) {
	var id string
	log.Printf("userService")
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		log.Printf("userService 1")
		id, errTx = s.userRepository.Create(ctx, info)
		log.Printf("userService 2")
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return id, nil
}
