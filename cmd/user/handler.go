package main

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/converter"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	repo := repository.NewRepository(repository.DB)
	userRepo := repository.NewUserRepository(repo)
	usr, err := userRepo.QueryUserByID(req.UserId)

	// 从数据库中查询到的user，转换为proto生成后的类型
	userResp, err := converter.ConvertUserModelToProto(usr)
	if err != nil {
		return nil, err
	}

	// TODO(century): 这里只是测试用，还需完善
	res := &user.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		User:       userResp,
	}

	return res, nil
}