package main

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/converter"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	repo := repository.NewRepository(repository.DB)
	userRepo := repository.NewUserRepository(repo)
	usr, err := userRepo.QueryUserByID(req.UserId)

	if err != nil {
		// TODO: 添加日志

		res := &user.UserInfoResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "query failed from database",
			User:       nil,
		}
		return res, err
	}

	// 从数据库中查询到的user，转换为proto生成后的类型
	userResp, err := converter.ConvertUserModelToProto(usr)
	if err != nil {
		return nil, err
	}

	// TODO(century): 这里只是测试用，还需完善
	res := &user.UserInfoResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "query user info success",
		User:       userResp,
	}

	return res, nil
}
