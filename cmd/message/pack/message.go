package pack

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/message"
)

// Note pack note info
func Message(m *model.Message) *message.Message {
	if m == nil {
		return nil
	}
	return &message.Message{
		Id:         int64(m.ID),
		ToUserId:   int64(m.ToUserID),
		FromUserId: int64(m.FromUserID),
		Content:    m.Content,
		CreateTime: m.CreatedAt.Unix(),
	}
}
func ChatList(cs []*model.Message) []*message.Message {
	chat := make([]*message.Message, 0)
	for _, c := range cs {
		if n := Message(c); n != nil {
			chat = append(chat, n)
		}
	}
	return chat
}
