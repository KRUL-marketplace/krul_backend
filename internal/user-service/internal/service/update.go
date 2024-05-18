package service

import (
	"context"
	"user-service/internal/repository/model"
)

func (s *userService) Update(ctx context.Context, id string, info *model.UserInfo) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.userRepository.Update(ctx, id, info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return err
}
