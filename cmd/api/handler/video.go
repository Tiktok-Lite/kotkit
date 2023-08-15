package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
	"time"
)

func Feed(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	latestTime := c.Query("latest_time")
	token := c.Query("token")
	var latestTime64 int64
	var parseErr error
	if latestTime != "" {
		latestTime64, parseErr = strconv.ParseInt(latestTime, 10, 64)
	} else {
		latestTime64 = time.Now().Unix()
	}

	if parseErr != nil {
		logger.Errorf("Failed to parse latest_time. %v", parseErr)
		ResponseError(c, http.StatusBadRequest, response.PackFeedError("请检查latest_time是否合法"))
	}
	// TODO(century): token后面处理

	req := &video.FeedRequest{
		LatestTime: &latestTime64,
		Token:      &token,
	}
	resp, err := rpc.Feed(ctx, req)
	if err != nil {
		logger.Errorf("RPC call failed due to %v", err)
		ResponseError(c, http.StatusInternalServerError, response.PackFeedError("由于内部错误，获取视频失败"))
	}

	ResponseSuccess(c, response.PackFeedSuccess(resp.NextTime, resp.VideoList, resp.StatusMsg))
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	userID := c.Query("user_id")
	token := c.Query("token")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		logger.Errorf("Failed to parse user_id. %v", err)
		ResponseError(c, http.StatusBadRequest, response.PackPublishListError("请检查user_id是否合法"))
		return
	}

	req := &video.PublishListRequest{
		UserId: id,
		Token:  token,
	}
	resp, err := rpc.PublishList(ctx, req)
	if err != nil {
		logger.Errorf("RPC call failed due to %v", err)
		ResponseError(c, http.StatusInternalServerError, response.PackPublishListError("由于内部错误，获取视频失败"))
		return
	}

	ResponseSuccess(c, response.PackPublishListSuccess(resp.VideoList, resp.StatusMsg))
}
