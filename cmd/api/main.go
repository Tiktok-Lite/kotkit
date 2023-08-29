package main

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/cmd/api/handler"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	v "github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var (
	logger     = log.Logger()
	apiConfig  = v.LoadConfig(constant.DefaultAPIConfigName)
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
			user.POST("/register/", handler.Register)
			user.POST("/login/", handler.Login)
		}
		publish := douyin.Group("/publish")
		{
			publish.GET("/list/", handler.PublishList)
			publish.POST("/action/", handler.PublishAction)
		}
		relation := douyin.Group("/relation")
		{
			// 粉丝列表
			relation.GET("/follower/list/", handler.FollowerList)
			// 关注列表
			relation.GET("/follow/list/", handler.FollowList)
			// 朋友列表
			relation.GET("/friend/list/", handler.FriendList)
			// 关注和取关
			relation.POST("/action/", handler.RelationAction)
		}
		douyin.GET("/feed/", handler.Feed)
		favorite := douyin.Group("/favorite")
		{
			favorite.POST("/action/", handler.FavoriteAction)
			favorite.GET("/list/", handler.FavoriteList)
		}
		comment := douyin.Group("/comment")
		{
			comment.GET("/list/", handler.CommentList)
			comment.POST("/action/", handler.CommentAction)
		}
		message := douyin.Group("/message")
		{
			message.GET("/chat/", handler.Chat)
			message.POST("/action/", handler.MessageAction)
		}
	}
}

func main() {
	// 初始化RPC客户端
	rpc.InitRPC()
	svr := server.Default(server.WithHostPorts(serverAddr))
	apiRegister(svr)

	svr.Spin()
	logger.Infof("HTTP service starts successfully at %s", serverAddr)
}
