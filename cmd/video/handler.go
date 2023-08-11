package main

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/converter"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (*video.FeedResponse, error) {
	nextTime := time.Now().Unix()
	repo := repository.NewRepository(repository.DB)
	videoRepo := repository.NewVideoRepository(repo)
	videos, err := videoRepo.Feed(req.LatestTime, req.Token)
	// 从倒序的videos中找到最小的nextTime
	if len(videos) != 0 {
		nextTime = videos[len(videos)-1].UpdatedAt.Unix()
	}

	if err != nil {
		return nil, err
	}
	videoListProto, err := converter.ConvertVideoModelListToProto(videos)
	if err != nil {
		return nil, err
	}

	res := &video.FeedResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "成功获取视频",
		VideoList:  videoListProto,
		NextTime:   &nextTime,
	}

	return res, nil
}
