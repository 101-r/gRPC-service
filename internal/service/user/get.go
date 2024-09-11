package user

import (
	"context"
	"log"

	"github.com/101-r/gRPC-service/internal/model"
)

func (s *service) Get(ctx context.Context, id int) (*model.UserInfo, error) {
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		log.Printf("error getting user: %v\n", err)
		return nil, err
	}

	if user == nil {
		log.Printf("error finding user: %v\n", err)
		return nil, model.ErrorUserNotFound
	}

	return user, nil
}
