package main

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

var (
	repo         = repository.NewRepository(db.DB())
	relationRepo = repository.NewRelationRepository(repo)
)

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	logger := log.Logger()
	claims, err := Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationActionResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	userID := claims.Id
	toUserID := req.ToUserId
	if userID == toUserID {
		logger.Errorf("操作非法：用户无法成为自己的粉丝：%d", userID)
		res := &relation.RelationActionResponse{
			StatusCode: -1,
			StatusMsg:  "操作非法：用户无法成为自己的粉丝",
		}
		return res, nil
	}
	if req.ActionType != 1 && req.ActionType != 2 {
		logger.Errorf("action_type 格式错误")
		res := &relation.RelationActionResponse{
			StatusCode: -1,
			StatusMsg:  "action_type 格式错误",
		}
		return res, nil
	}
	if err != nil {
		logger.Errorf("Error occurs when converting video lists to proto. %v", err)
		return nil, err
	}

	res := &relation.RelationActionResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "操作成功",
	}

	return res, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	logger := log.Logger()
	userID := req.UserId

	// 解析token,获取用户id
	claims, err := Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFollowListResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	if userID != claims.Id {
		logger.Errorf("当前登录用户%d无法访问其他用户的关注列表%d", claims.Id, userID)
		res := &relation.RelationFollowListResponse{
			StatusCode: -1,
			StatusMsg:  "当前登录用户无法访问其他用户的关注列表",
		}
		return res, nil
	}

	// 从数据库获取关注列表
	followings, err := relationRepo.GetFollowingListByUserID(uint(userID))
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFollowListResponse{
			StatusCode: -1,
			StatusMsg:  "关注列表获取失败",
		}
		return res, nil
	}
	userIDs := make([]int64, 0)
	for _, res := range followings {
		userIDs = append(userIDs, int64(res.ToUserId))
	}
	userList := make([]*user.User, 0)
	if err != nil {
		logger.Errorf("Error occurs when converting video lists to proto. %v", err)
		return nil, err
	}

	// 返回结果
	res := &relation.RelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   userList,
	}
	return res, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	// TODO: Your code here...
	logger := log.Logger()
	userID := req.UserId

	// 解析token,获取用户id
	claims, err := Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFollowerListResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	if userID != claims.Id {
		logger.Errorf("当前登录用户%d无法访问其他用户的粉丝列表%d", claims.Id, userID)
		res := &relation.RelationFollowerListResponse{
			StatusCode: -1,
			StatusMsg:  "当前登录用户无法访问其他用户的粉丝列表",
		}
		return res, nil
	}

	// 从数据库获取粉丝列表
	followers, err := relationRepo.GetFollowerListByUserID(uint(userID))
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFollowerListResponse{
			StatusCode: -1,
			StatusMsg:  "粉丝列表获取失败",
		}
		return res, nil
	}
	userIDs := make([]int64, 0)
	for _, res := range followers {
		userIDs = append(userIDs, int64(res.ToUserId))
	}
	userList := make([]*user.User, 0)
	if err != nil {
		logger.Errorf("Error occurs when converting video lists to proto. %v", err)
		return nil, err
	}

	// 返回结果
	res := &relation.RelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   userList,
	}
	return res, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
