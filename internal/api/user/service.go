package user

import (
	"github.com/101-r/gRPC-service/internal/service"

	desc "github.com/101-r/gRPC-service/pkg/user"
)

type Implementation struct {
	desc.UnimplementedUserServiceServer
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
