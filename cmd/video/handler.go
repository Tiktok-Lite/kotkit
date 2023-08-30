package main

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/video"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/converter"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/Tiktok-Lite/kotkit/pkg/oss"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

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

	videos, err := db.Feed(req.LatestTime)
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

	// 处理videos中的点赞关系和minio中的视频url和封面url
	for _, v := range videos {
		liked, err := db.QueryVideoLikeRelation(int64(v.ID), uid)
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

		playUrl, err := oss.GetObjectURL(oss.VideoBucketName, v.PlayURL)
		if err != nil {
			logger.Errorf("Error occurs when getting video url from minio. %v", err)
			res := &video.FeedResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "内部minio数据库错误，获取视频失败",
				VideoList:  nil,
				NextTime:   nil,
			}
			return res, nil
		}
		v.PlayURL = playUrl

		coverUrl, err := oss.GetObjectURL(oss.CoverBucketName, v.CoverURL)
		if err != nil {
			logger.Errorf("Error occurs when getting cover url from minio. %v", err)
			res := &video.FeedResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "内部minio数据库错误，获取视频失败",
				VideoList:  nil,
				NextTime:   nil,
			}
			return res, nil
		}
		v.CoverURL = coverUrl

		// 处理视频的用户是否被关注
		followed, err := db.QueryUserByRelation(int64(v.Author.ID), uid)
		v.Author.IsFollow = followed

		// 处理发布视频的用户的头像和背景图
		avatarUrl, err := oss.GetObjectURL(oss.AvatarBucketName, v.Author.Avatar)
		if err != nil {
			logger.Errorf("Error occurs when getting avatar url from minio. %v", err)
			res := &video.FeedResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "内部minio数据库错误，获取视频失败",
				VideoList:  nil,
				NextTime:   nil,
			}
			return res, nil
		}
		v.Author.Avatar = avatarUrl

		bgUrl, err := oss.GetObjectURL(oss.BackgroundImageBucketName, v.Author.BackgroundImage)
		if err != nil {
			logger.Errorf("Error occurs when getting background image url from minio. %v", err)
			res := &video.FeedResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "内部minio数据库错误，获取视频失败",
				VideoList:  nil,
				NextTime:   nil,
			}
			return res, nil
		}
		v.Author.BackgroundImage = bgUrl
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

	claims, err := Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("Error occurs when parsing token. %v", err)
		res := &video.PublishListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "token解析失败",
			VideoList:  nil,
		}
		return res, nil
	}

	videos, err := db.QueryVideoListByUserID(req.UserId)
	if err != nil {
		logger.Errorf("Error occurs when querying video list from database. %v", err)
		res := &video.PublishListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部数据库错误，获取视频失败",
			VideoList:  nil,
		}
		return res, nil
	}

	if videos == nil {
		res := &video.PublishListResponse{
			StatusCode: constant.StatusOKCode,
			StatusMsg:  fmt.Sprintf("用户id为%d的用户视频不存在", req.UserId),
			VideoList:  nil,
		}
		return res, nil
	}

	// 处理videos中的点赞关系和minio中的视频url和封面url
	for _, v := range videos {
		liked, err := db.QueryVideoLikeRelation(int64(v.ID), claims.Id)
		if err != nil {
			logger.Errorf("Error occurs when querying video like relation from database. %v", err)
			res := &video.PublishListResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "内部数据库错误，获取视频失败",
				VideoList:  nil,
			}
			return res, nil
		}
		v.IsFavorite = liked

		playUrl, err := oss.GetObjectURL(oss.VideoBucketName, v.PlayURL)
		if err != nil {
			logger.Errorf("Error occurs when getting video url from minio. %v", err)
			res := &video.PublishListResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "内部minio数据库错误，获取视频失败",
				VideoList:  nil,
			}
			return res, nil
		}
		v.PlayURL = playUrl

		coverUrl, err := oss.GetObjectURL(oss.CoverBucketName, v.CoverURL)
		if err != nil {
			logger.Errorf("Error occurs when getting cover url from minio. %v", err)
			res := &video.PublishListResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "内部minio数据库错误，获取视频失败",
				VideoList:  nil,
			}
			return res, nil
		}
		v.CoverURL = coverUrl
	}

	videoListProto, err := converter.ConvertVideoModelListToProto(videos)
	if err != nil {
		logger.Errorf("Error occurs when converting video lists to proto. %v", err)
		res := &video.PublishListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部转换错误，获取视频失败",
			VideoList:  nil,
		}
		return res, nil
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

	// 默认文件不得大于50MB
	if len(req.Data) > 50*1024*1024 {
		logger.Errorf("Video file is too large %v.", len(req.Data))
		res := &video.PublicActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "视频文件太大",
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

	// 注意保证数据的原子性
	err = db.CreateVideo(video_)
	if err != nil {
		logger.Errorf("Error occurs when creating video to database. %v", err)
		res := &video.PublicActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "创建视频失败",
		}
		return res, nil
	}

	go func() {
		err = oss.PublishVideo(req.Data, videoTitle, coverTitle)
		if err != nil {
			logger.Errorf("Error occurs when publishing video to minio. %v", err)
			// minio发布失败则删除数据库中的记录
			e := db.DeleteVideoById(video_.ID, uint(userID))
			if e != nil {
				logger.Errorf("Error occurs when deleting video and records. %v", e)
			}
		}
	}()

	res := &video.PublicActionResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "创建视频成功",
	}

	return res, nil
}
