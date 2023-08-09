package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

func UserInfo(ctx context.Context, c *app.RequestContext) {
	userID := c.Query("user_id")
	token := c.Query("token")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		// 先不处理，后面再做
		return
	}
	req := &user.UserInfoRequest{
		UserId: id,
		Token:  token,
	}
	resp, _ := rpc.UserInfo(ctx, req)
	if resp.StatusCode == constant.StatusOKCode {
		c.JSON(http.StatusOK, response.UserInfo{
			Base: response.Base{
				StatusCode: constant.StatusOKCode,
				StatusMsg:  resp.StatusMsg,
			},
			User: resp.User,
		})
		return
	}

	c.JSON(http.StatusOK, response.UserInfo{
		Base: response.Base{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "用户信息获取失败",
		},
		User: nil,
	})
}
