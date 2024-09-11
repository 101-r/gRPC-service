package converter

import (
	"github.com/101-r/gRPC-service/internal/model"

	repoModel "github.com/101-r/gRPC-service/internal/repository/model"
)

func ToUserInfoFromRepo(info *repoModel.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Id:        int64(info.Id),
		Username:  info.Username,
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Email:     info.Email,
		Password:  info.Password,
		CreatedAt: info.CreatedAt,
		UpdatedAt: &info.UpdatedAt.Time,
	}
}
