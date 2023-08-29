package rpc

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/comment"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/comment/commentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	commentClient commentservice.Client
)

func InitComment(config *viper.Viper) {
	commentServiceName := config.GetString("server.name")
	commentServiceAddr := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))

	c, err := commentservice.NewClient(commentServiceName, client.WithHostPorts(commentServiceAddr))
	if err != nil {
		panic(err)
	}

	commentClient = c
}

func CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (*comment.DouyinCommentListResponse, error) {
	return commentClient.CommentList(ctx, req)
}
func CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (*comment.DouyinCommentActionResponse, error) {
	return commentClient.CommentAction(ctx, req)
}
