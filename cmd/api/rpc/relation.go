package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation/relationservice"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	relationClient relationservice.Client
)

func InitRelation(config *viper.Viper) {
	relationServiceName := config.GetString("server.name")
	relationServiceAddr := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))

	c, err := relationservice.NewClient(relationServiceName, client.WithHostPorts(relationServiceAddr))
	if err != nil {
		panic(err)
	}

	relationClient = c
}

func RelationActionList(ctx context.Context, req *relation.RelationActionRequest) (*relation.RelationActionResponse, error) {
	return relationClient.RelationAction(ctx, req)
}

func RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (*relation.RelationFollowListResponse, error) {
	return relationClient.RelationFollowList(ctx, req)
}

func RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (*relation.RelationFollowerListResponse, error) {
	return relationClient.RelationFollowerList(ctx, req)
}

func RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (*relation.RelationFriendListResponse, error) {
	return relationClient.RelationFriendList(ctx, req)
}
