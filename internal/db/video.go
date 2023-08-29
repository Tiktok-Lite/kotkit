package db

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

func Feed(latestTime *int64) ([]*model.Video, error) {
	logger := log.Logger()
	// 不指定latest time则根据当前时间指定
	if latestTime == nil || *latestTime == 0 {
		curr := time.Now().UnixMilli()
		latestTime = &curr
	}
	var videos []*model.Video

	// 注意：Preload内的参数是字段名，而不是表名......
	if err := DB().Debug().
		Preload("Author").Where("created_at < ?", time.UnixMilli(*latestTime)).Order("created_at desc").
		Limit(30).Find(&videos).Error; err != nil {
		logger.Errorf("failed to query videos from databse: %v", err)
		return nil, err
	}

	return videos, nil
}

func QueryVideoListByUserID(userID int64) ([]*model.Video, error) {
	logger := log.Logger()

	var videos []*model.Video

	if err := DB().Debug().Preload("Author").Where("user_id = ?", userID).Find(&videos).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		logger.Errorf("failed to query videos from databse: %v", err)
		return nil, err
	}

	return videos, nil
}

func CreateVideo(video *model.Video) error {
	if err := DB().Debug().Create(video).Error; err != nil {
		return err
	}
	return nil

}

func QueryVideoLikeRelation(vid, uid int64) (bool, error) {
	var count int64
	if err := DB().Debug().Raw("SELECT COUNT(*) FROM user_like_videos WHERE video_id = ? AND user_id = ?", vid, uid).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func AddVideoFavoriteCountById(vid int64, tx *gorm.DB) error {
	var video model.Video
	if err := tx.Debug().Model(&video).Where("id = ?", vid).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func DeleteVideoFavoriteCountById(vid int64, tx *gorm.DB) error {
	var video model.Video
	if err := tx.Debug().Model(&video).Where("id = ?", vid).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func QueryUserIdByVideoId(vid int64) (uint, error) {
	var video model.Video
	if err := DB().Debug().Where("id = ?", vid).First(&video).Error; err != nil {
		return 0, err
	}
	return video.UserID, nil
}

func QueryFavoriteVideoIdsByUserId(uid int64) ([]int64, error) {
	var videoIds []int64
	err := DB().Debug().Raw("SELECT video_id FROM user_like_videos WHERE user_id = ?", uid).Scan(&videoIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return videoIds, nil
}

func QueryVideosByVideoIds(videoIds []int64) ([]*model.Video, error) {
	var videos []*model.Video
	err := DB().Debug().Preload("Author").Where("id IN ?", videoIds).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
