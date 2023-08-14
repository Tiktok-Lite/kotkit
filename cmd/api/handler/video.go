package handler

import (
	"bytes"
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"io"
	"net/http"
	"strconv"
	"time"
)

var logger = log.Logger()

func Feed(ctx context.Context, c *app.RequestContext) {
	latestTime := c.Query("latest_time")
	token := c.Query("token")
	var latestTime64 int64
	if latestTime != "" {
		latestTime64, _ = strconv.ParseInt(latestTime, 10, 64)
	} else {
		latestTime64 = time.Now().Unix()
	}
	// TODO(century): token后面处理

	req := &video.FeedRequest{
		LatestTime: &latestTime64,
		Token:      &token,
	}
	resp, err := rpc.Feed(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, response.Feed{
			Base: response.Base{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "RPC调用出现问题",
			},
			NextTime:  nil,
			VideoList: nil,
		})
	}
	if resp.StatusCode == constant.StatusOKCode {
		c.JSON(http.StatusOK, response.Feed{
			Base: response.Base{
				StatusCode: constant.StatusOKCode,
				StatusMsg:  "成功获取feed流",
			},
			NextTime:  resp.NextTime,
			VideoList: resp.VideoList,
		})
	}
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	userID := c.Query("user_id")
	token := c.Query("token")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		logger.Errorf("Failed to parse user_id. %v", err)
		c.JSON(http.StatusBadRequest, response.PublishList{
			Base: response.Base{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "用户ID非法!",
			},
			VideoList: nil,
		})
	}

	req := &video.PublishListRequest{
		UserId: id,
		Token:  token,
	}
	resp, err := rpc.PublishList(ctx, req)
	if err != nil {
		logger.Errorf("RPC call failed due to %v", err)
		c.JSON(http.StatusInternalServerError, response.PublishList{
			Base: response.Base{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "由于内部错误，获取视频失败",
			},
			VideoList: nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.PublishList{
		Base: response.Base{
			StatusCode: constant.StatusOKCode,
			StatusMsg:  resp.StatusMsg,
		},
		VideoList: resp.VideoList,
	})
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
