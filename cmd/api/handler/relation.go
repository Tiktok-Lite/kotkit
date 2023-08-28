package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

func FollowerList(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	userID := c.Query("user_id")
	if userID == "" {
		logger.Info("Input user_id is empty")
		ResponseError(c, http.StatusBadRequest, response.PackListError("user_id不能为空"))
		return
	}
	token := c.Query("token")
	if token == "" {
		logger.Info("Input token is empty")
		ResponseError(c, http.StatusBadRequest, response.PackListError("token不能为空"))
		return
	}

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse user_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackListError("粉丝列表获取失败，请检查输入是否合法"))
		return
	}
	req := &relation.RelationFollowerListRequest{
		UserId: id,
		Token:  token,
	}
	resp, err := rpc.RelationFollowerList(ctx, req)
	if err != nil {
		logger.Errorf("failed to call rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError, response.Relation{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  resp.StatusMsg,
			},
		})
	}
	ResponseSuccess(c, response.PackListSuccess(resp.UserList, "粉丝列表获取成功"))
}

func FollowList(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	userID := c.Query("user_id")
	if userID == "" {
		logger.Info("Input user_id is empty")
		ResponseError(c, http.StatusBadRequest, response.PackListError("user_id不能为空"))
		return
	}
	token := c.Query("token")
	if token == "" {
		logger.Info("Input token is empty")
		ResponseError(c, http.StatusBadRequest, response.PackListError("token不能为空"))
		return
	}
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse user_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackListError("关注列表获取失败，请检查输入是否合法"))
		return
	}
	req := &relation.RelationFollowListRequest{
		UserId: id,
		Token:  token,
	}
	resp, err := rpc.RelationFollowList(ctx, req)
	if err != nil {
		logger.Errorf("failed to call rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError,
			response.PackListError(resp.StatusMsg))
		return
	}
	ResponseSuccess(c, response.PackListSuccess(resp.UserList, "关注列表获取成功"))
}

func FriendList(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	userID := c.Query("user_id")
	if userID == "" {
		logger.Info("Input user_id is empty")
		ResponseError(c, http.StatusBadRequest, response.PackListError("user_id不能为空"))
		return
	}
	token := c.Query("token")
	if token == "" {
		logger.Info("Input token is empty")
		ResponseError(c, http.StatusBadRequest, response.PackListError("token不能为空"))
		return
	}
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse user_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackListError("好友列表获取失败，请检查输入是否合法"))
		return
	}
	req := &relation.RelationFriendListRequest{
		UserId: id,
		Token:  token,
	}
	resp, err := rpc.RelationFriendList(ctx, req)
	if err != nil {
		logger.Errorf("failed to call rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError,
			response.PackListError(resp.StatusMsg))
		return
	}
	ResponseSuccess(c, response.PackListSuccess(resp.UserList, "好友列表获取成功"))
}

func RelationAction(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()
	toUserID := c.Query("to_user_id")
	if toUserID == "" {
		logger.Info("Input to_user_id is empty")
		ResponseError(c, http.StatusBadRequest, response.PackActionError("to_user_id不能为空"))
		return
	}
	tid, err := strconv.ParseInt(toUserID, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse user_id: %v", err)
		ResponseError(c, http.StatusOK, response.Relation{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "to_user_id 不合法",
			},
		})
		return
	}
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil || (actionType != 1 && actionType != 2) {
		ResponseError(c, http.StatusOK, response.Relation{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "action_type 不合法",
			},
		})
		return
	}
	token := c.Query("token")
	if token == "" {
		logger.Info("Input token is empty")
		ResponseError(c, http.StatusBadRequest, response.PackPublishListError("token不能为空"))
		return
	}
	req := &relation.RelationActionRequest{
		ToUserId:   tid,
		Token:      token,
		ActionType: int32(actionType),
	}
	resp, err := rpc.RelationAction(ctx, req)
	if err != nil {
		logger.Errorf("failed to call rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError,
			response.PackActionError(resp.StatusMsg))
		return
	}
	ResponseSuccess(c, response.PackActionSuccess(resp.StatusMsg))
}
