package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video/videoservice"
	"github.com/Tiktok-Lite/kotkit/pkg/etcd"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	videoClient videoservice.Client
)

func InitVideo(config *viper.Viper) {
	r, err := etcd.Resolver()
	if err != nil {
		logger.Errorf("Error occurs when creating etcd resolver: %v", err)
		panic(err)
	}

	videoServiceName := config.GetString("server.name")

	c, err := videoservice.NewClient(videoServiceName, client.WithResolver(r))
	if err != nil {
		logger.Errorf("Error occurs when creating video client: %v", err)
		panic(err)
	}

	videoClient = c
}

func Feed(ctx context.Context, req *video.FeedRequest) (*video.FeedResponse, error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}
	return resp, nil
}

func PublishList(ctx context.Context, req *video.PublishListRequest) (*video.PublishListResponse, error) {
	resp, err := videoClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}
	return resp, nil
}

func PublishAction(ctx context.Context, req *video.PublishActionRequest) (*video.PublicActionResponse, error) {
	resp, err := videoClient.PublishAction(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}
