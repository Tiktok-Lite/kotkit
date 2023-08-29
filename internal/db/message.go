package db

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/pkg/errors"
)

func SendMessage(message *model.Message) error {
	if err := db.Save(message).Error; err != nil {
		return errors.New("failed to send message")
	}
	return nil
}
func QueryMessageList(userId int64, toUserId int64) ([]*model.Message, error) {
	var messages []*model.Message
	if err := db.Debug().Where("from_user_id", userId).Where("to_user_id", toUserId).
		Order("created_at desc").Find(&messages).Error; err != nil {
		logger.Errorf("failed to query messages from databse: %v", err)
		return nil, err
	}
	return messages, nil
}
