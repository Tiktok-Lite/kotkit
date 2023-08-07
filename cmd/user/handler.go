package main

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	repo := repository.NewRepository(repository.DB)
	userRepo := repository.NewUserRepository(repo)
	usr, err := userRepo.QueryUserByID(req.UserId)

	// TODO(century): 这里只是测试用，还需完善
	res := &user.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		User: &user.User{
			Id:              usr.ID,
			Name:            usr.Name,
			FollowCount:     nil,
			FollowerCount:   nil,
			IsFollow:        usr.IsFollow,
			Avatar:          &usr.Avatar,
			BackgroundImage: &usr.BackgroundImage,
			Signature:       &usr.Signature,
			TotalFavorited:  nil,
			WorkCount:       nil,
			FavoriteCount:   nil,
		},
	}

	return res, nil
}
