package main

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/cmd/api/handler"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	v "github.com/Tiktok-Lite/kotkit/pkg/viper"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var (
	apiConfig  = v.LoadConfig(constant.DefaultAPIConfigPath)
	serverAddr = fmt.Sprintf("%s:%d", apiConfig.GetString("server.host"), apiConfig.GetInt("server.port"))
)

func apiRegister(hz *server.Hertz) {
	// 连通性测试，后续完成接口开发后会删掉
	hz.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	douyin := hz.Group("/douyin")
	{
		user := douyin.Group("/user")
		{
			user.GET("/", handler.UserInfo)
		}
	}
}

func main() {
	// 初始化RPC客户端
	rpc.InitRPC()
	svr := server.Default(server.WithHostPorts(serverAddr))
	apiRegister(svr)

	svr.Spin()
}
