package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

func Feed(ctx context.Context, c *app.RequestContext) {
	latestTime := c.Query("latest_time")
	token := c.Query("token")
	t, err := strconv.ParseInt(latestTime, 10, 64)
	if err != nil {
		return
	}

	req := &video.FeedRequest{
		LatestTime: &t,
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
