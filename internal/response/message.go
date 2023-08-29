package response

import (
	"github.com/Tiktok-Lite/kotkit/kitex_gen/message"
)

type MessageList struct {
	Base
	MessageList []*message.Message `json:"message"`
}

type Message struct {
	Base
	Message *message.Message `json:"message"`
}

func PackMessageListSuccess(messageList []*message.Message, msg string) MessageList {
	base := PackBaseSuccess(msg)
	return MessageList{
		Base:        base,
		MessageList: messageList,
	}
}

func PackMessageActionSuccess(msg string) Message {
	base := PackBaseSuccess(msg)
	return Message{
		Base: base,
	}
}
