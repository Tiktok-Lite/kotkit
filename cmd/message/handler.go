package main

import (
	"context"
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

	res := &message.DouyinMessageChatResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "success",
	}
	token := req.Token
	parseToken, err := Jwt.ParseToken(token)
	if err != nil {
		logger.Errorf("Error occurs when parsing token. %v", err)
		res.StatusMsg = "token错误"
		return res, err
	}
	userId := parseToken.Id
	chat, err := db.QueryMessageList(userId, req.ToUserId)
	if err != nil {
		logger.Errorf("Error occurs when querying chat list from database. %v", err)
		return nil, err
	}
	chatList := pack.ChatList(chat)
	res.MessageList = chatList
	return res, nil
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	logger := log.Logger()

	res := &message.DouyinMessageActionResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "success",
	}
	claims, err := Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("Error occurs when parsing token. %v", err)
		res.StatusMsg = "token错误"
		return res, err
	}

	if req.ActionType == 1 {
		c := model.Message{
			Content:    req.Content,
			ToUserID:   uint(req.ToUserId),
			FromUserID: uint(claims.Id),
			CreateTime: new(time.Time).Format("01-02"),
		}
		err := db.SendMessage(&c)
		if err != nil {
			logger.Errorf("Error occurs when add message to database. %v", err)
			res.StatusMsg = "发送消息失败"
		}
		return res, err
	} else {
		res.StatusCode = constant.StatusErrorCode
		res.StatusMsg = "ActionType错误"
		return res, err
	}
}
