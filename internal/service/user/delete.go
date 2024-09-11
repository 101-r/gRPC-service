package user

import (
	"context"
	"log"
)

func (s *service) Delete(ctx context.Context, id int) (int, error) {
	id, err := s.userRepository.Delete(ctx, id)
	if err != nil {
		log.Printf("error deliting user: %v\n", err)
		return -1, err
	}

	return id, nil
}
