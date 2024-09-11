package app

import (
	"log"

	"github.com/101-r/gRPC-service/internal/api/user"
	"github.com/101-r/gRPC-service/internal/config"
	"github.com/101-r/gRPC-service/internal/repository"
	"github.com/101-r/gRPC-service/internal/service"

	userRepository "github.com/101-r/gRPC-service/internal/repository/user"
	userService "github.com/101-r/gRPC-service/internal/service/user"
)

type serviceProvider struct {
	config config.Config

	userRepository repository.UserRepository

	userService service.UserService

	userImpl *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("failed to load config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		connStr := s.Config().ConnStr()
		s.userRepository = userRepository.NewRepository(connStr)
	}

	return s.userRepository
}

func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(),
		)
	}

	return s.userService
}

func (s *serviceProvider) UserImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(
			s.UserService(),
		)
	}

	return s.userImpl
}
