package main

import (
	"context"
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
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	userID := claims.Id
	toUserID := req.ToUserId
	if userID == toUserID {
		logger.Errorf("操作非法：用户无法成为自己的粉丝：%d", userID)
		res := &relation.RelationActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "操作非法：用户无法成为自己的粉丝",
		}
		return res, nil
	}
	err = db.RelationAction(claims.Id, req)
	if err != nil {
		logger.Errorf("数据库内部错误：%v", err)
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
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	if userID != claims.Id {
		logger.Errorf("当前登录用户%d无法访问其他用户的关注列表%d", claims.Id, userID)
		res := &relation.RelationFollowListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "当前登录用户无法访问其他用户的关注列表",
		}
		return res, nil
	}
	//从数据库获取关注列表
	followings, err := db.GetFollowingListByUserID(uint(userID))
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFollowListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "关注列表获取失败",
		}
		return res, nil
	}
	userListProto, err := converter.ConvertFollowingListModelToProto(followings)
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFollowListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部转换错误，获取关注列表失败",
			UserList:   nil,
		}
		return res, nil
	}

	// 返回结果
	res := &relation.RelationFollowListResponse{
		StatusCode: constant.StatusOKCode,
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
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	if userID != claims.Id {
		logger.Errorf("当前登录用户%d无法访问其他用户的粉丝列表%d", claims.Id, userID)
		res := &relation.RelationFollowerListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "当前登录用户无法访问其他用户的粉丝列表",
		}
		return res, nil
	}
	// 从数据库获取粉丝列表
	followers, err := db.GetFollowerListByUserID(uint(userID))
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFollowerListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "粉丝列表获取失败",
		}
		return res, nil
	}
	userListProto, err := converter.ConvertFollowerListModelToProto(followers)
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFollowerListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部转换错误，获取粉丝列表失败",
			UserList:   nil,
		}
		return res, nil
	}

	// 返回结果
	res := &relation.RelationFollowerListResponse{
		StatusCode: constant.StatusOKCode,
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
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	if userID != claims.Id {
		logger.Errorf("当前登录用户%d无法访问其他用户的朋友列表%d", claims.Id, userID)
		res := &relation.RelationFriendListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "当前登录用户无法访问其他用户的朋友列表",
		}
		return res, nil
	}
	// 从数据库获取朋友列表
	friends, err := db.GetFriendList(uint(userID))
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFriendListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "好友列表获取失败",
		}
		return res, nil
	}
	userListProto, err := converter.ConvertFollowerListModelToProto(friends)
	if err != nil {
		logger.Errorf(err.Error())
		res := &relation.RelationFriendListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部转换错误，获取朋友列表失败",
			UserList:   nil,
		}
		return res, nil
	}

	// 返回结果
	res := &relation.RelationFriendListResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "success",
		UserList:   userListProto,
	}
	return res, nil
}
