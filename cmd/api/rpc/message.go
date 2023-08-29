package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/message"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/message/messageservice"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	messageClient messageservice.Client
)

func InitMessage(config *viper.Viper) {
	messageServiceName := config.GetString("server.name")
	messageServiceAddr := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))

	c, err := messageservice.NewClient(messageServiceName, client.WithHostPorts(messageServiceAddr))
	if err != nil {
		panic(err)
	}

	messageClient = c
}

func MessageList(ctx context.Context, req *message.DouyinMessageChatRequest) (*message.DouyinMessageChatResponse, error) {
	return messageClient.MessageChat(ctx, req)
}
func MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (*message.DouyinMessageActionResponse, error) {
	return messageClient.MessageAction(ctx, req)
}
