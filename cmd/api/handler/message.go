package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/message"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
	"strings"
)

func Chat(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	token := c.Query("token")
	if token == "" {
		logger.Errorf("Illegal input: empty token.")
		ResponseError(c, http.StatusBadRequest, response.PackMessageListError("token不能为空"))
		return
	}

	toUserId := c.Query("to_user_id")
	if toUserId == "" {
		logger.Errorf("Illegal input: empty to_user_id.")
		ResponseError(c, http.StatusBadRequest, response.PackMessageListError("to_user_id不能为空"))
		return
	}
	id, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse to_user_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackMessageListError("请检查您的输入是否合法"))
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
			response.PackMessageListError(resp.StatusMsg))
		return
	}

	ResponseSuccess(c, response.PackMessageListSuccess(resp.MessageList, "聊天列表获取成功"))
}

func MessageAction(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	token := c.Query("token")
	if token == "" {
		logger.Errorf("Illegal input: empty token.")
		ResponseError(c, http.StatusBadRequest, response.PackMessageListError("token不能为空"))
		return
	}

	toUserId := c.Query("to_user_id")
	if toUserId == "" {
		logger.Errorf("Illegal input: empty to_user_id.")
		ResponseError(c, http.StatusBadRequest, response.PackMessageListError("to_user_id不能为空"))
		return
	}
	id, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse to_user_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackMessageListError("请检查您的输入是否合法"))
		return
	}

	actionType := c.Query("action_type")
	if actionType == "" {
		logger.Errorf("Illegal input: empty action_type.")
		ResponseError(c, http.StatusBadRequest, response.PackMessageListError("action_type不能为空"))
		return
	}
	act, err := strconv.Atoi(actionType)
	if err != nil {
		logger.Errorf("failed to parse action_type: %v", err)
		ResponseError(c, http.StatusBadRequest, response.PackMessageListError("请检查您的输入是否合法"))
		return
	}

	req := &message.DouyinMessageActionRequest{
		Token:      token,
		ToUserId:   id,
		ActionType: int32(act),
	}
	if act == constant.PostMessageCode {
		content := c.Query("content")
		s := strings.TrimSpace(content)
		if len(s) == 0 {
			ResponseError(c, http.StatusInternalServerError,
				response.PackMessageListError("消息不能为空"))
			return
		}
		req.Content = s
	}
	resp, err := rpc.MessageAction(ctx, req)
	if err != nil {
		logger.Errorf("failed to call rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError,
			response.PackMessageListError(resp.StatusMsg))
		return
	}

	ResponseSuccess(c, response.PackMessageActionSuccess(resp.StatusMsg))
}
