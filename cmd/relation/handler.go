package main

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/cmd/relation/command"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/converter"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

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
	err = command.NewRelationActionService(ctx).RelationAction(claims.Id, req)
	if err != nil {
		logger.Errorf("Error occurs when converting relation lists to proto. %v", err)
		return nil, err
	}

	if req.ActionType == 1 {
		res := &relation.RelationActionResponse{
			StatusCode: constant.StatusOKCode,
			StatusMsg:  "关注成功",
		}
		return res, nil
	}

	if req.ActionType == 2 {
		res := &relation.RelationActionResponse{
			StatusCode: constant.StatusOKCode,
			StatusMsg:  "取关成功",
		}
		return res, nil
	}

	return nil, nil
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
	//从数据库获取关注列表
	followings, err := db.GetFollowingListByUserID(uint(userID))
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFollowListResponse{
			StatusCode: -1,
			StatusMsg:  "关注列表获取失败",
		}
		return res, nil
	}
	userListProto, err := converter.ConvertFollowingListModelToProto(followings)
	fmt.Println(userListProto)
	if err != nil {
		return nil, err
	}

	// 返回结果
	res := &relation.RelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   userListProto,
	}
	return res, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
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
	followers, err := db.GetFollowerListByUserID(uint(userID))
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFollowerListResponse{
			StatusCode: -1,
			StatusMsg:  "粉丝列表获取失败",
		}
		return res, nil
	}
	userListProto, err := converter.ConvertFollowerListModelToProto(followers)
	if err != nil {
		return nil, err
	}

	// 返回结果
	res := &relation.RelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   userListProto,
	}
	return res, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	logger := log.Logger()
	userID := req.UserId

	// 解析token,获取用户id
	claims, err := Jwt.ParseToken(req.Token)

	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFriendListResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	if userID != claims.Id {
		logger.Errorf("当前登录用户%d无法访问其他用户的朋友列表%d", claims.Id, userID)
		res := &relation.RelationFriendListResponse{
			StatusCode: -1,
			StatusMsg:  "当前登录用户无法访问其他用户的朋友列表",
		}
		fmt.Println(res)
		return res, nil
	}
	// 从数据库获取朋友列表
	friends, err := db.GetFriendList(uint(userID))
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFriendListResponse{
			StatusCode: -1,
			StatusMsg:  "好友列表获取失败",
		}
		return res, nil
	}
	userListProto, err := converter.ConvertFollowerListModelToProto(friends)
	if err != nil {
		return nil, err
	}

	// 返回结果
	res := &relation.RelationFriendListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   userListProto,
	}
	return res, nil
}
