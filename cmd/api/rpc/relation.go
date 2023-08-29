package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation/relationservice"
	"github.com/Tiktok-Lite/kotkit/pkg/etcd"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	relationClient relationservice.Client
)

func InitRelation(config *viper.Viper) {
	r, err := etcd.Resolver()
	if err != nil {
		logger.Errorf("Error occurs when creating etcd resolver: %v", err)
		panic(err)
	}

	relationServiceName := config.GetString("server.name")

	c, err := relationservice.NewClient(relationServiceName, client.WithResolver(r))
	if err != nil {
		logger.Errorf("Error occurs when creating relation client: %v", err)
		panic(err)
	}

	relationClient = c
}

func RelationAction(ctx context.Context, req *relation.RelationActionRequest) (*relation.RelationActionResponse, error) {
	resp, err := relationClient.RelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}
	return resp, nil
}

func RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (*relation.RelationFollowListResponse, error) {
	resp, err := relationClient.RelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}
	return resp, nil
}

func RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (*relation.RelationFollowerListResponse, error) {
	resp, err := relationClient.RelationFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}
	return resp, nil
}

func RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (*relation.RelationFriendListResponse, error) {
	resp, err := relationClient.RelationFriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}
	return resp, nil
}
