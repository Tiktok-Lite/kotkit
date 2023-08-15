package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Register(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	username := c.Query("username")
	password := c.Query("password")
	//校验参数
	if len(username) == 0 || len(password) == 0 {
		ResponseError(c, http.StatusBadRequest, response.PackLoginOrRegisterError("用户名或密码不能为空"))
		return
	}
	if len(username) > 32 || len(password) > 32 {
		ResponseError(c, http.StatusBadRequest, response.PackLoginOrRegisterError("用户名或密码长度不能大于32个字符"))
		return
	}

	req := &login.UserRegisterRequest{
		Username: username,
		Password: password,
	}
	resp, err := rpc.Register(ctx, req)
	if err != nil {
		logger.Errorf("register error: %v", err)
		ResponseError(c, http.StatusInternalServerError, response.PackLoginOrRegisterError("服务器内部错误，注册失败"))
		return
	}

	ResponseSuccess(c, response.PackLoginOrRegisterSuccess(resp.UserId, resp.Token, resp.StatusMsg))
}

func Login(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		ResponseError(c, http.StatusBadRequest, response.PackLoginOrRegisterError("用户名或密码不能为空"))
		return
	}

	req := &login.UserLoginRequest{
		Username: username,
		Password: password,
	}
	resp, err := rpc.Login(ctx, req)
	if err != nil {
		logger.Errorf("login error: %v", err)
		ResponseError(c, http.StatusInternalServerError, response.PackLoginOrRegisterError("服务器内部错误，登录失败"))
		return
	}
	
	ResponseSuccess(c, response.PackLoginOrRegisterSuccess(resp.UserId, resp.Token, resp.StatusMsg))
}
