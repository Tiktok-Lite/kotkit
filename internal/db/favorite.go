package db

import (
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"gorm.io/gorm"
)

func AddLikeVideoRelation(vid, uid int64, tx *gorm.DB) error {
	err := tx.Exec("INSERT INTO user_like_videos (video_id, user_id) VALUES (?, ?)", vid, uid).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func DeleteLikeVideoRelation(vid, uid int64, tx *gorm.DB) error {
	err := tx.Exec("DELETE FROM user_like_videos WHERE video_id = ? AND user_id = ?", vid, uid).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// AddLikeVideo
// 当点赞时：
// 1. user_like_videos表中插入一条记录
// 2. video表中的favorite_count字段+1
// 3. 点赞user表中的favorite_count字段+1
// 4. 拥有video的user表中的total_favorited字段+1
func AddLikeVideo(vid, uid, ownerId int64) error {
	logger := log.Logger()

	tx := DB().Begin()
	err := AddLikeVideoRelation(vid, uid, tx)
	if err != nil {
		logger.Errorf("Failed to add like video relation due to %v", err)
		return err
	}

	err = AddVideoFavoriteCountById(vid, tx)
	if err != nil {
		logger.Errorf("Failed to add video favorite count due to %v", err)
		return err
	}

	err = AddUserFavoriteCountVideoById(uid, tx)
	if err != nil {
		logger.Errorf("Failed to add user favorite count due to %v", err)
		return err
	}

	err = AddUserTotalFavoritedById(ownerId, tx)
	if err != nil {
		logger.Errorf("Failed to add user total favorited due to %v", err)
		return err
	}

	tx.Commit()
	return nil
}

// DislikeVideo
// 当取消点赞时：
// 1. user_like_videos表中删除一条记录
// 2. video表中的favorite_count字段-1
// 3. 点赞user表中的favorite_count字段-1
// 4. 拥有video的user表中的total_favorited字段-1
func DislikeVideo(vid, uid, ownerId int64) error {
	tx := DB().Begin()
	err := DeleteLikeVideoRelation(vid, uid, tx)
	if err != nil {
		return err
	}

	err = DeleteVideoFavoriteCountById(vid, tx)
	if err != nil {
		return err
	}

	err = DeleteUserFavoriteCountVideoById(uid, tx)
	if err != nil {
		return err
	}

	err = DeleteUserTotalFavoritedById(ownerId, tx)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
