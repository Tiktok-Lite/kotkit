package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/message"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/message/messageservice"
	"github.com/Tiktok-Lite/kotkit/pkg/etcd"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	messageClient messageservice.Client
)

func InitMessage(config *viper.Viper) {
	r, err := etcd.Resolver()
	if err != nil {
		logger.Errorf("Error occurs when creating etcd resolver: %v", err)
		panic(err)
	}

	messageServiceName := config.GetString("server.name")

	c, err := messageservice.NewClient(messageServiceName, client.WithResolver(r))

	if err != nil {
		logger.Errorf("Error occurs when creating login client: %v", err)
		panic(err)
	}
	messageClient = c
}

func MessageList(ctx context.Context, req *message.DouyinMessageChatRequest) (*message.DouyinMessageChatResponse, error) {
	resp, err := messageClient.MessageChat(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}
	return resp, nil
}
func MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (*message.DouyinMessageActionResponse, error) {
	resp, err := messageClient.MessageAction(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode == constant.StatusErrorCode {
		return resp, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}
