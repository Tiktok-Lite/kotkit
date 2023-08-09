package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	//校验参数
	if len(username) == 0 || len(password) == 0 {
		c.JSON(http.StatusBadRequest, login.UserRegisterResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "用户名或密码不能为空",
		})
		return
	}
	
	req := &login.UserRegisterRequest{
		Username: username,
		Password: password,
	}
	resp, err := rpc.Register(ctx, req)
	if err != nil {
		// 处理错误 TODO
		return
	}
	
	if resp.StatusCode == -1 {
		c.JSON(http.StatusOK, login.UserRegisterResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	c.JSON(http.StatusOK, login.UserRegisterResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  resp.StatusMsg,
		UserID:     resp.UserID,
		Token:		resp.Token
	})

}

func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	
	if len(username) == 0 || len(password) == 0 {
		c.JSON(http.StatusBadRequest, login.UserLoginResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "用户名或密码不能为空",
		})
		return
	}
	
	req := &login.UserLoginRequest{
		Username: username,
		Password: password,
	}
	resp, err := rpc.Login(ctx, req)
	if err != nil {
		// 处理错误 TODO
		return
	}
	if resp.StatusCode == -1 {
		c.JSON(http.StatusOK, login.UserLoginResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	c.JSON(http.StatusOK, login.UserLoginResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  resp.StatusMsg,
		UserID:     resp.UserID,
		Token:		resp.Token
	})
}
