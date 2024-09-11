package user

import (
	"context"
	"log"

	"github.com/101-r/gRPC-service/internal/converter"

	desc "github.com/101-r/gRPC-service/pkg/user"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &desc.CreateResponse{
		Id: int64(id),
	}, nil
}
