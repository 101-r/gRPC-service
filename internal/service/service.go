package service

import (
	"context"

	"github.com/101-r/gRPC-service/internal/model"
)

type UserService interface {
	Create(ctx context.Context, info *model.UserInfo) (int, error)
	Get(ctx context.Context, id int) (*model.UserInfo, error)
	Delete(ctx context.Context, id int) (int, error)
}
