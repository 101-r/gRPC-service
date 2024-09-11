package user

import (
	"context"
	"log"

	"github.com/101-r/gRPC-service/internal/converter"
	desc "github.com/101-r/gRPC-service/pkg/user"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, int(req.Id))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &desc.GetResponse{
		User: converter.ToUser(user),
	}, nil
}
