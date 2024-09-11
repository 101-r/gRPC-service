package user

import (
	"github.com/101-r/gRPC-service/internal/repository"

	def "github.com/101-r/gRPC-service/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

func NewService(
	userRepository repository.UserRepository,
) *service {
	return &service{
		userRepository: userRepository,
	}
}
