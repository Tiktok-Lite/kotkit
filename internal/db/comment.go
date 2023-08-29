package db

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/pkg/errors"
)

func AddComment(comment *model.Comment) error {
	if err := db.Save(comment).Error; err != nil {
		return errors.New("failed to add comment")
	}
	return nil
}

func QueryCommentByVideoID(videoId int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	if err := db.Debug().Where("video_id", videoId).
		Order("created_at desc").Find(&comments).Error; err != nil {
		logger.Errorf("failed to query comments from databse: %v", err)
		return nil, err
	}
	return comments, nil
}

func DeleteCommentById(commentId int64) error {
	if err := db.Delete(&model.Comment{}, commentId).Error; err != nil {
		return errors.New("failed to delete comment")
	}
	return nil
}
