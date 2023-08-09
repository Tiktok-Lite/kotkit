package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/login/userservice"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	v "github.com/Tiktok-Lite/kotkit/pkg/viper"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	userClient userservice.Client
	userConfig = v.LoadConfig(constant.DefaultUserConfigPath)
)

func InitUser(config *viper.Config) {
	userServiceName := config.GetString("server.name")
	userServiceAddr := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))
	
	c, err := userservice.NewClient(userServiceName, client.WithHostPorts(userServiceAddr))
	
	if err != nil {
		panic(err)
	}
	userClient = c
}

func Register(ctx context.Context, req *login.UserRegisterRequest) (*login.UserRegisterResponse, error) {
	return userClient.Register(ctx, req)
}

func Login(ctx context.Context, req *login.UserLoginRequest) (*login.UserLoginResponse, error) {
	return userClient.Login(ctx, req)
}

