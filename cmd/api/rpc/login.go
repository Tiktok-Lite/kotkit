package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login/loginservice"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	loginClient loginservice.Client
)

func InitLogin(config *viper.Viper) {
	loginServiceName := config.GetString("server.name")
	loginServiceAddr := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))

	c, err := loginservice.NewClient(loginServiceName, client.WithHostPorts(loginServiceAddr))

	if err != nil {
		panic(err)
	}
	loginClient = c
}

func Register(ctx context.Context, req *login.UserRegisterRequest) (*login.UserRegisterResponse, error) {
	resp, err := loginClient.Register(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}
	return resp, nil
}

func Login(ctx context.Context, req *login.UserLoginRequest) (*login.UserLoginResponse, error) {
	resp, err := loginClient.Login(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}
	return resp, nil
}
