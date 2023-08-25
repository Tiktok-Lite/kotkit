package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video/videoservice"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	videoClient videoservice.Client
)

func InitVideo(config *viper.Viper) {
	videoServiceName := config.GetString("server.name")
	videoServiceAddr := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))

	c, err := videoservice.NewClient(videoServiceName, client.WithHostPorts(videoServiceAddr))
	if err != nil {
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
	return videoClient.PublishList(ctx, req)
}

func PublishAction(ctx context.Context, req *video.PublishActionRequest) (*video.PublicActionResponse, error) {
	resp, err := videoClient.PublishAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}
