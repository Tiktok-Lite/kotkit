package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	userClient userservice.Client
)

func InitUser(config *viper.Viper) {
	userServiceName := config.GetString("server.name")
	userServiceAddr := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))

	c, err := userservice.NewClient(userServiceName, client.WithHostPorts(userServiceAddr))
	if err != nil {
		panic(err)
	}

	userClient = c
}

func UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	return userClient.UserInfo(ctx, req)
}
