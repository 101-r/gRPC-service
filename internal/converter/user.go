package converter

import (
	"github.com/101-r/gRPC-service/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/101-r/gRPC-service/pkg/user"
)

func ToUser(info *model.UserInfo) *desc.User {
	info_ := desc.UserInfo{
		Id:        info.Id,
		Username:  info.Username,
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Email:     info.Email,
		Password:  info.Password,
		CreatedAt: timestamppb.New(info.CreatedAt),
		UpdatedAt: timestamppb.New(*info.UpdatedAt),
	}

	return &desc.User{
		Id:   info.Id,
		Info: &info_,
	}
}

func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Username:  info.Username,
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Email:     info.Email,
		Password:  info.Password,
	}
}
