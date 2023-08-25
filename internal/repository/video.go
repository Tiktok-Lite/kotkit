package repository

import (
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"time"
)

type VideoRepository interface {
	Feed(latestTime *int64) ([]*model.Video, error)
	QueryVideoListByUserID(userID int64, token string) ([]*model.Video, error)
	CreateVideo(video *model.Video) error
	QueryVideoLikeRelation(vid, uid int64) (bool, error)
}

type videoRepository struct {
	*Repository
}

func NewVideoRepository(r *Repository) VideoRepository {
	return &videoRepository{
		Repository: r,
	}
}

func (v *videoRepository) Feed(latestTime *int64) ([]*model.Video, error) {
	// 不指定latest time则根据当前时间指定
	curr := time.Now().UnixMilli()
	if latestTime == nil {
		latestTime = &curr
	}
	var videos []*model.Video

	// 注意：Preload内的参数是字段名，而不是表名......
	if err := v.db.Debug().
		Preload("Author").Where("created_at < ?", time.UnixMilli(*latestTime)).Order("created_at desc").
		Find(&videos).Error; err != nil {
		v.logger.Errorf("failed to query videos from databse: %v", err)
		return nil, err
	}

	return videos, nil
}

func (v *videoRepository) QueryVideoListByUserID(userID int64, token string) ([]*model.Video, error) {
	var videos []*model.Video

	if err := v.db.Debug().Preload("Author").Where("user_id = ?", userID).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}

func (v *videoRepository) CreateVideo(video *model.Video) error {
	if err := v.db.Debug().Create(video).Error; err != nil {
		return err
	}
	return nil

}

func (v *videoRepository) QueryVideoLikeRelation(vid, uid int64) (bool, error) {
	var count int64
	if err := db.DB().Debug().Raw("SELECT COUNT(*) FROM user_like_videos WHERE video_id = ? AND user_id = ?", vid, uid).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
