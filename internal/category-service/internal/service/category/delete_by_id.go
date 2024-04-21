package category

import "context"

func (s *serv) DeleteById(ctx context.Context, id uint32) error {
	err := s.categoryRepository.DeleteById(ctx, id)

	return err
}
