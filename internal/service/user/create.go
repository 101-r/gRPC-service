package user

import (
	"context"
	"log"

	"github.com/101-r/gRPC-service/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.UserInfo) (int, error) {
	id, err := s.userRepository.Create(ctx, info)
	if err != nil {
		log.Printf("error creating user: %v\n", err)
		return -1, err
	}

	return id, nil
}
