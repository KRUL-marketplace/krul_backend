package product

import "context"

func (s *serv) DeleteById(ctx context.Context, id string) error {
	err := s.productRepository.DeleteById(ctx, id)

	return err
}
