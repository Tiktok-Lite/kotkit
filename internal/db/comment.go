package db

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func AddComment(comment *model.Comment, tx *gorm.DB) error {
	if err := tx.Save(comment).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to add comment")
	}
	return nil
}

func QueryCommentByVideoID(videoId int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	if err := DB().Debug().Where("video_id", videoId).
		Order("created_at desc").Find(&comments).Error; err != nil {
		logger.Errorf("failed to query comments from databse: %v", err)
		return nil, err
	}
	return comments, nil
}

func DeleteCommentById(commentId int64, tx *gorm.DB) error {
	if err := tx.Unscoped().Delete(&model.Comment{}, commentId).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete comment")
	}
	return nil
}

func CommentTransaction(comment *model.Comment) error {
	tx := DB().Begin()
	if err := AddComment(comment, tx); err != nil {
		return err
	}
	if err := AddVideoCommentCountById(comment.VideoID, tx); err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func UnCommentTransaction(commentId int64, vid uint) error {
	tx := DB().Begin()
	if err := DeleteCommentById(commentId, tx); err != nil {
		return err
	}
	if err := DeleteVideoCommentCountById(vid, tx); err != nil {
		return err
	}

	tx.Commit()
	return nil
}
