package service

import "context"

func (s *userService) DeleteById(ctx context.Context, id string) error {
	err := s.userRepository.DeleteById(ctx, id)

	return err
}
