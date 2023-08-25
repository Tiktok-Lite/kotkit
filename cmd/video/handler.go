package main

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/converter"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/Tiktok-Lite/kotkit/pkg/oss"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

var (
	repo      = repository.NewRepository(db.DB())
	videoRepo = repository.NewVideoRepository(repo)
)

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (*video.FeedResponse, error) {
	logger := log.Logger()
	nextTime := time.Now().UnixMilli()

	var uid int64
	if len(*req.Token) > 0 {
		if claims, err := Jwt.ParseToken(*req.Token); err != nil {
			logger.Errorf("Error occurs when parsing token. %v", err)
			res := &video.FeedResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "token解析失败",
				VideoList:  nil,
				NextTime:   nil,
			}
			return res, nil
		} else {
			uid = claims.Id
		}
	}

	videos, err := videoRepo.Feed(req.LatestTime)
	if err != nil {
		logger.Errorf("Error occurs when querying video list from database. %v", err)
		res := &video.FeedResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部数据库错误，获取视频失败",
			VideoList:  nil,
			NextTime:   nil,
		}
		return res, nil
	}

	// 从倒序的videos中找到最小的nextTime
	if len(videos) != 0 {
		nextTime = videos[len(videos)-1].UpdatedAt.Unix()
	}

	// 处理videos中的点赞关系
	for _, v := range videos {
		liked, err := videoRepo.QueryVideoLikeRelation(int64(v.ID), uid)
		if err != nil {
			logger.Errorf("Error occurs when querying video like relation from database. %v", err)
			res := &video.FeedResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "内部数据库错误，获取视频失败",
				VideoList:  nil,
				NextTime:   nil,
			}
			return res, nil
		}
		v.IsFavorite = liked
	}

	videoListProto, err := converter.ConvertVideoModelListToProto(videos)
	if err != nil {
		logger.Errorf("Error occurs when converting video lists to proto. %v", err)
		res := &video.FeedResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部转换错误，获取视频失败",
			VideoList:  nil,
			NextTime:   nil,
		}
		return res, nil
	}

	res := &video.FeedResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "成功获取视频",
		VideoList:  videoListProto,
		NextTime:   &nextTime,
	}
	return res, nil
}

func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (*video.PublishListResponse, error) {
	logger := log.Logger()

	videos, err := videoRepo.QueryVideoListByUserID(req.UserId, req.Token)
	if err != nil {
		logger.Errorf("Error occurs when querying video list from database. %v", err)
		return nil, err
	}
	videoListProto, err := converter.ConvertVideoModelListToProto(videos)
	if err != nil {
		logger.Errorf("Error occurs when converting video lists to proto. %v", err)
		return nil, err
	}

	res := &video.PublishListResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "成功获取视频",
		VideoList:  videoListProto,
	}

	return res, nil
}

func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (*video.PublicActionResponse, error) {
	logger := log.Logger()

	claims, err := Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("Error occurs when parsing token. %v", err)
		res := &video.PublicActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "token解析失败",
		}
		return res, nil
	}

	if len(req.Title) == 0 || len(req.Title) > 32 {
		logger.Errorf("Title is empty or too long.")
		res := &video.PublicActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "标题不能为空或者太长",
		}
		return res, nil
	}

	userID := claims.Id
	videoTitle, coverTitle := fmt.Sprintf("%d_%s_%d.mp4", userID, req.Title, time.Now().Unix()), fmt.Sprintf("%d_%s_%d.jpg", userID, req.Title, time.Now().Unix())

	video_ := &model.Video{
		UserID:   uint(userID),
		PlayURL:  videoTitle,
		CoverURL: coverTitle,
		Title:    req.Title,
	}

	err = oss.PublishVideo(req.Data, videoTitle, coverTitle)
	if err != nil {
		logger.Errorf("Error occurs when publishing video. %v", err)
		res := &video.PublicActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "上传视频到minio失败",
		}
		return res, nil
	}

	// TODO: 在上传数据库和OSS中间加入事务
	// 注意：先把东西上传到oss后写入数据库，目的是防止上传失败后数据库中有记录但是oss中没有
	// 保证数据的原子性
	err = videoRepo.CreateVideo(video_)
	if err != nil {
		logger.Errorf("Error occurs when creating video to database. %v", err)
		res := &video.PublicActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "创建视频失败",
		}
		return res, nil
	}

	res := &video.PublicActionResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "创建视频成功",
	}

	return res, nil
}
