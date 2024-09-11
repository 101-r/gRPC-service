package user

import (
	"context"
	"log"

	desc "github.com/101-r/gRPC-service/pkg/user"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	id, err := i.userService.Delete(ctx, int(req.Id))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &desc.DeleteResponse{
		Id: int64(id),
	}, nil
}
