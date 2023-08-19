package handler

import (
	"bytes"
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"io"
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

func PublishAction(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	token := c.PostForm("token")
	if token == "" {
		logger.Info("Input token is empty")
		c.JSON(http.StatusBadRequest, response.PublishAction{
			Base: response.Base{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "token不能为空",
			},
		})
		return
	}

	title := c.PostForm("title")
	if title == "" {
		logger.Info("Input title is empty")
		c.JSON(http.StatusBadRequest, response.PublishAction{
			Base: response.Base{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "title不能为空",
			},
		})
		return
	}

	videoFile, err := c.FormFile("data")
	if err != nil {
		logger.Errorf("Failed to get video file from request. %v", err)
		c.JSON(http.StatusBadRequest, response.PublishAction{
			Base: response.Base{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "获取视频文件失败",
			},
		})
		return
	}
	src, err := videoFile.Open()
	defer src.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		logger.Errorf("Failed to read video file. %v", err)
		c.JSON(http.StatusBadRequest, response.PublishAction{
			Base: response.Base{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "读取视频文件失败",
			},
		})
		return
	}

	req := &video.PublishActionRequest{
		Token: token,
		Data:  buf.Bytes(),
		Title: title,
	}
	res, err := rpc.PublishAction(ctx, req)
	if err != nil {
		logger.Errorf("RPC call failed due to %v", err)
		c.JSON(http.StatusInternalServerError, response.PublishAction{
			Base: response.Base{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "由于内部错误，发布视频失败",
			},
		})
		return
	}
	c.JSON(http.StatusOK, response.PublishAction{
		Base: response.Base{
			StatusCode: constant.StatusOKCode,
			StatusMsg:  res.StatusMsg,
		},
	})
}
