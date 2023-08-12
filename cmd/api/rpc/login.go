package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login/loginservice"
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
	return loginClient.Register(ctx, req)
}

func Login(ctx context.Context, req *login.UserLoginRequest) (*login.UserLoginResponse, error) {
	return loginClient.Login(ctx, req)
}

