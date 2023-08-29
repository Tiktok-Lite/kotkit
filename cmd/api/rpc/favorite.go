package rpc

import (
	"context"
	"errors"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/favorite"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/favorite/favoriteservice"
	"github.com/Tiktok-Lite/kotkit/pkg/etcd"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	favoriteClient favoriteservice.Client
)

func InitFavorite(config *viper.Viper) {
	r, err := etcd.Resolver()
	if err != nil {
		logger.Errorf("Error occurs when creating etcd resolver: %v", err)
		panic(err)
	}

	favoriteServiceName := config.GetString("server.name")

	c, err := favoriteservice.NewClient(favoriteServiceName, client.WithResolver(r))
	if err != nil {
		logger.Errorf("Error occurs when creating favorite client: %v", err)
		panic(err)
	}

	favoriteClient = c
}

func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, errors.New(resp.StatusMsg)
	}

	return resp, nil
}

func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	resp, err := favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, errors.New(*resp.StatusMsg)
	}

	return resp, nil
}
