package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/message"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
	"strings"
)

func Chat(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	toUserId := c.Query("to_user_id")
	token := c.Query("token")
	id, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse video_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackBaseError("请检查您的输入是否合法"))
		return
	}
	req := &message.DouyinMessageChatRequest{
		Token:    token,
		ToUserId: id,
	}
	resp, err := rpc.MessageList(ctx, req)
	if err != nil {
		logger.Errorf("failed to call rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError,
			response.PackBaseError("聊天列表获取失败，服务器内部问题"))
		return
	}

	ResponseSuccess(c, response.PackMessageListSuccess(resp.MessageList, "聊天列表获取成功"))
}

func MessageAction(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	toUserId := c.Query("to_user_id")
	id, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse to_user_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackBaseError("请检查您的输入是否合法"))
		return
	}

	actionType := c.Query("action_type")
	act, err := strconv.Atoi(actionType)
	if err != nil {
		logger.Errorf("failed to parse action_type: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackBaseError("请检查您的输入是否合法"))
		return
	}
	token := c.Query("token")

	req := &message.DouyinMessageActionRequest{
		Token:      token,
		ToUserId:   id,
		ActionType: int32(act),
	}
	if act == 1 {
		content := c.Query("content")
		s := strings.TrimSpace(content)
		if len(s) == 0 {
			ResponseError(c, http.StatusInternalServerError,
				response.PackBaseError("消息不能为空"))
			return
		}
		req.Content = s
	}

	_, err = rpc.MessageAction(ctx, req)
	if err != nil {
		logger.Errorf("failed to call rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError,
			response.PackBaseError("发送消息失败，服务器内部问题"))
		return
	}

	ResponseSuccess(c, response.PackMessageActionSuccess("发送成功"))
}
