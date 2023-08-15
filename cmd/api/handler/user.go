package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

func UserInfo(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	userID := c.Query("user_id")
	token := c.Query("token")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse user_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackUserInfoError("请检查您的输入是否合法"))
		return
	}
	req := &user.UserInfoRequest{
		UserId: id,
		Token:  token,
	}
	resp, err := rpc.UserInfo(ctx, req)
	if err != nil {
		logger.Errorf("failed to call rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError,
			response.PackUserInfoError("用户信息获取失败，服务器内部问题"))
		return
	}

	ResponseSuccess(c, response.PackUserInfoSuccess(resp.User, "用户信息获取成功"))
}
