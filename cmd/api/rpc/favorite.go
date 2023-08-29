package rpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/favorite"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/favorite/favoriteservice"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	favoriteClient favoriteservice.Client
)

func InitFavorite(config *viper.Viper) {
	favoriteServiceName := config.GetString("server.name")
	favoriteServiceAddr := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))

	c, err := favoriteservice.NewClient(favoriteServiceName, client.WithHostPorts(favoriteServiceAddr))
	if err != nil {
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
