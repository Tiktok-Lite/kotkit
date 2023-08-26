package main

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/converter"
	"github.com/Tiktok-Lite/kotkit/pkg/oss"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	repo := repository.NewRepository(db.DB())
	userRepo := repository.NewUserRepository(repo)

	token := req.Token
	claims, err := Jwt.ParseToken(token)
	if err != nil {
		logger.Errorf("Failed to authenticate due to %v", err)
		res := &user.UserInfoResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "鉴权失败，请检查您的token合法性",
			User:       nil,
		}
		return res, nil
	}

	usr, err := userRepo.QueryUserByID(req.UserId)

	if err == nil && usr == nil {
		logger.Errorf("No user exists due to user_id: %v", req.UserId)
		res := &user.UserInfoResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "用户不存在",
			User:       nil,
		}
		return res, nil
	} else if err != nil {
		logger.Errorf("Failed to query from database due to %v", err)
		res := &user.UserInfoResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部数据库查询错误",
			User:       nil,
		}
		return res, nil
	}

	isRelated, err := userRepo.QueryUserByRelation(req.UserId, claims.Id)
	if err != nil {
		logger.Errorf("Failed to query from database due to %v", err)
		res := &user.UserInfoResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部数据库查询错误",
			User:       nil,
		}
		return res, nil
	}
	// 如果查询不到，说明当前用户没有关注该用户
	usr.IsFollow = isRelated

	avatarUrl, err := oss.GetObjectURL(oss.AvatarBucketName, usr.Avatar)
	if err != nil {
		logger.Errorf("Failed to get object url due to %v", err)
		res := &user.UserInfoResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "minio数据库查询错误",
			User:       nil,
		}
		return res, nil
	}
	usr.Avatar = avatarUrl

	bgImgUrl, err := oss.GetObjectURL(oss.BackgroundImageBucketName, usr.BackgroundImage)
	if err != nil {
		logger.Errorf("Failed to get object url due to %v", err)
		res := &user.UserInfoResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "minio数据库查询错误",
			User:       nil,
		}
		return res, nil
	}
	usr.BackgroundImage = bgImgUrl

	// 从数据库中查询到的user，转换为proto生成后的类型
	userResp, err := converter.ConvertUserModelToProto(usr)
	if err != nil {
		logger.Errorf("Failed to convert user due to %v", err)
		res := &user.UserInfoResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部转换出现错误",
			User:       nil,
		}
		return res, nil
	}

	res := &user.UserInfoResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "查询用户成功",
		User:       userResp,
	}

	return res, nil
}
