package repository

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"time"
)

type VideoRepository interface {
	Feed(latestTime *int64, token *string) ([]*model.Video, error)
	QueryVideoListByUserID(userID int64, token string) ([]*model.Video, error)
}

type videoRepository struct {
	*Repository
}

func NewVideoRepository(r *Repository) VideoRepository {
	return &videoRepository{
		Repository: r,
	}
}

func (v *videoRepository) Feed(latestTime *int64, token *string) ([]*model.Video, error) {
	// TODO(century): 处理token，目前简化处理
	// 不指定latest time则根据当前时间指定
	t := time.Now()
	if latestTime != nil {
		t = time.Unix(*latestTime, 0)
	}
	var videos []*model.Video

	// 注意：Preload内的参数是字段名，而不是表名......
	if err := v.db.Debug().
		Preload("Author").Where("created_at < ?", t).Order("created_at desc").
		Find(&videos).Error; err != nil {
		v.logger.Errorf("failed to query videos from databse: %v", err)
		return nil, err
	}

	return videos, nil
}

func (v *videoRepository) QueryVideoListByUserID(userID int64, token string) ([]*model.Video, error) {
	// TODO(century): 处理token，目前简化处理
	var videos []*model.Video

	if err := v.db.Debug().Preload("Author").Where("user_id = ?", userID).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}
