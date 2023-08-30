package main

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/cmd/message/pack"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/message"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"time"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	parseToken, err := Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("Error occurs when parsing token. %v", err)
		res := &message.DouyinMessageChatResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}

	res := &message.DouyinMessageChatResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "success",
	}
	userId := parseToken.Id
	// 自己的聊天列表
	userChat, err := db.QueryMessageList(userId, req.ToUserId)
	// 对方的聊天列表
	toUserChat, err := db.QueryMessageList(req.ToUserId, userId)
	if err != nil {
		logger.Errorf("Error occurs when querying chat list from database. %v", err)
		res.StatusCode = constant.StatusErrorCode
		res.StatusMsg = "查询消息列表失败"
		return res, nil
	}
	// 合并聊天列表
	chatList := pack.ChatList(userChat)
	chatList = append(chatList, pack.ChatList(toUserChat)...)
	res.MessageList = chatList
	return res, nil
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	logger := log.Logger()

	parseToken, err := Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("Error occurs when parsing token. %v", err)
		res := &message.DouyinMessageActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	res := &message.DouyinMessageActionResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "success",
	}

	if req.ActionType == constant.PostMessageCode {
		t := time.Now()
		c := model.Message{
			Content:    req.Content,
			ToUserID:   uint(req.ToUserId),
			FromUserID: uint(parseToken.Id),
			CreateTime: fmt.Sprintf("%02d-%02d", t.Month(), t.Day()),
		}
		err := db.SendMessage(&c)
		if err != nil {
			logger.Errorf("Error occurs when add message to database. %v", err)
			res.StatusCode = constant.StatusErrorCode
			res.StatusMsg = "发送消息失败"
			return res, nil
		}
		return res, nil
	}

	res.StatusCode = constant.StatusErrorCode
	res.StatusMsg = "ActionType错误"
	return res, nil
}
