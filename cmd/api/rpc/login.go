package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login/loginservice"
	"github.com/Tiktok-Lite/kotkit/pkg/etcd"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	loginClient loginservice.Client
)

func InitLogin(config *viper.Viper) {
	r, err := etcd.Resolver()
	if err != nil {
		logger.Errorf("Error occurs when creating etcd resolver: %v", err)
		panic(err)
	}

	loginServiceName := config.GetString("server.name")

	c, err := loginservice.NewClient(loginServiceName, client.WithResolver(r))

	if err != nil {
		logger.Errorf("Error occurs when creating login client: %v", err)
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
