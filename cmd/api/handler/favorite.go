package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/favorite"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	if token == "" {
		logger.Errorf("Illegal input: empty token.")
		ResponseError(c, http.StatusBadRequest, response.PackFavoriteActionError("token不能为空"))
		return
	}

	if videoId == "" {
		logger.Errorf("Illegal input: empty video_id.")
		ResponseError(c, http.StatusBadRequest, response.PackFavoriteActionError("video_id不能为空"))
		return
	}

	vid, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse video_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackFavoriteActionError("请检查您的输入是否合法"))
		return
	}

	if actionType == "" {
		logger.Errorf("Illegal input: empty action_type.")
		ResponseError(c, http.StatusBadRequest, response.PackFavoriteActionError("action_type不能为空"))
		return
	}
	action, err := strconv.ParseInt(actionType, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse action_type: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackFavoriteActionError("请检查您的输入是否合法"))
		return
	}

	req := &favorite.FavoriteActionRequest{
		Token:      token,
		VideoId:    vid,
		ActionType: int32(action),
	}
	resp, err := rpc.FavoriteAction(ctx, req)
	if err != nil {
		logger.Errorf("error occurs when calling rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError, response.PackFavoriteActionError(resp.StatusMsg))
		return
	}

	ResponseSuccess(c, response.PackFavoriteActionSuccess(resp.StatusMsg))
}

func FavoriteList(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	token := c.Query("token")
	userId := c.Query("user_id")

	if token == "" {
		logger.Errorf("Illegal input: empty token.")
		ResponseError(c, http.StatusBadRequest, response.PackFavoriteListError("token不能为空"))
		return
	}

	if userId == "" {
		logger.Errorf("Illegal input: empty user_id.")
		ResponseError(c, http.StatusBadRequest, response.PackFavoriteListError("user_id不能为空"))
		return
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse user_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackFavoriteListError("请检查您的输入是否合法"))
		return
	}

	req := &favorite.FavoriteListRequest{
		UserId: uid,
		Token:  token,
	}
	resp, err := rpc.FavoriteList(ctx, req)
	if err != nil {
		logger.Errorf("error occurs when calling rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError, response.PackFavoriteListError(*resp.StatusMsg))
		return
	}

	ResponseSuccess(c, response.PackFavoriteListSuccess(resp.VideoList, *resp.StatusMsg))
}
