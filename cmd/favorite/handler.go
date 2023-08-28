package main

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	favorite "github.com/Tiktok-Lite/kotkit/kitex_gen/favorite"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/converter"
	"github.com/Tiktok-Lite/kotkit/pkg/oss"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
	token := req.Token
	claims, err := Jwt.ParseToken(token)
	if err != nil {
		logger.Errorf("Failed to authenticate due to %v", err)
		res := &favorite.FavoriteActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "鉴权失败，请检查您的token合法性",
		}
		return res, nil
	}

	if req.ActionType != constant.FavoriteCode && req.ActionType != constant.UnFavoriteCode {
		logger.Errorf("Failed to like video due to invalid action type")
		res := &favorite.FavoriteActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "无效的操作类型",
		}
		return res, nil
	}

	ownerId, err := db.QueryUserIdByVideoId(req.VideoId)
	if err != nil {
		logger.Errorf("Failed to like video due to %v", err)
		res := &favorite.FavoriteActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部数据库错误",
		}
		return res, nil
	}

	// 点赞操作
	if req.ActionType == constant.FavoriteCode {
		err = db.AddLikeVideo(req.VideoId, claims.Id, int64(ownerId))
		if err != nil {
			logger.Errorf("Failed to like video due to %v", err)
			res := &favorite.FavoriteActionResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "内部数据库错误",
			}
			return res, nil
		}
		res := &favorite.FavoriteActionResponse{
			StatusCode: constant.StatusOKCode,
			StatusMsg:  "点赞成功",
		}
		return res, nil
	}

	// 取消点赞操作
	err = db.DislikeVideo(req.VideoId, claims.Id, int64(ownerId))
	if err != nil {
		logger.Errorf("Failed to like video due to %v", err)
		res := &favorite.FavoriteActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "内部数据库错误",
		}
		return res, nil
	}

	res := &favorite.FavoriteActionResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "取消点赞成功",
	}
	return res, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	token := req.Token
	claims, err := Jwt.ParseToken(token)
	if err != nil {
		logger.Errorf("Failed to authenticate due to %v", err)
		errorMsg := "鉴权失败，请检查您的token合法性"
		res := &favorite.FavoriteListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  &errorMsg,
		}
		return res, nil
	}

	ids, err := db.QueryFavoriteVideoIdsByUserId(req.UserId)
	if err == nil && ids == nil {
		msg := "用户没有喜欢的视频"
		res := &favorite.FavoriteListResponse{
			StatusCode: constant.StatusOKCode,
			StatusMsg:  &msg,
			VideoList:  nil,
		}
		return res, nil
	}
	if err != nil {
		logger.Errorf("Failed to query favorite video ids due to %v", err)
		errorMsg := "内部数据库错误"
		res := &favorite.FavoriteListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  &errorMsg,
		}
		return res, nil
	}

	videos, err := db.QueryVideosByVideoIds(ids)
	if err == nil && videos == nil {
		msg := "用户没有喜欢的视频"
		res := &favorite.FavoriteListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  &msg,
			VideoList:  nil,
		}
		return res, nil
	}
	if err != nil {
		logger.Errorf("Failed to query videos by video ids due to %v", err)
		errorMsg := "内部数据库错误"
		res := &favorite.FavoriteListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  &errorMsg,
		}
		return res, nil
	}

	// 处理videos中的点赞关系和minio中的视频url和封面url
	for _, v := range videos {
		liked, err := db.QueryVideoLikeRelation(int64(v.ID), claims.Id)
		if err != nil {
			logger.Errorf("Error occurs when querying video like relation from database. %v", err)
			errorMsg := "内部数据库错误，获取视频失败"
			res := &favorite.FavoriteListResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  &errorMsg,
				VideoList:  nil,
			}
			return res, nil
		}
		v.IsFavorite = liked

		// 获取author头像和背景图
		avatar, err := oss.GetObjectURL(oss.AvatarBucketName, v.Author.Avatar)
		if err != nil {
			logger.Errorf("Error occurs when getting avatar url from minio. %v", err)
			errorMsg := "内部minio数据库错误，获取author头像失败"
			res := &favorite.FavoriteListResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  &errorMsg,
				VideoList:  nil,
			}
			return res, nil
		}
		v.Author.Avatar = avatar

		bgImg, err := oss.GetObjectURL(oss.BackgroundImageBucketName, v.Author.BackgroundImage)
		if err != nil {
			logger.Errorf("Error occurs when getting background image url from minio. %v", err)
			errorMsg := "内部minio数据库错误，获取author背景图失败"
			res := &favorite.FavoriteListResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  &errorMsg,
				VideoList:  nil,
			}
			return res, nil
		}
		v.Author.BackgroundImage = bgImg

		playUrl, err := oss.GetObjectURL(oss.VideoBucketName, v.PlayURL)
		if err != nil {
			logger.Errorf("Error occurs when getting video url from minio. %v", err)
			errorMsg := "内部minio数据库错误，获取视频失败"
			res := &favorite.FavoriteListResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  &errorMsg,
				VideoList:  nil,
			}
			return res, nil
		}
		v.PlayURL = playUrl

		coverUrl, err := oss.GetObjectURL(oss.CoverBucketName, v.CoverURL)
		if err != nil {
			logger.Errorf("Error occurs when getting cover url from minio. %v", err)
			errorMsg := "内部minio数据库错误，获取视频失败"
			res := &favorite.FavoriteListResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  &errorMsg,
				VideoList:  nil,
			}
			return res, nil
		}
		v.CoverURL = coverUrl
	}

	videoListProto, err := converter.ConvertVideoModelListToProto(videos)
	if err != nil {
		logger.Errorf("Error occurs when converting video lists to proto. %v", err)
		errorMsg := "内部转换错误，获取视频失败"
		res := &favorite.FavoriteListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  &errorMsg,
			VideoList:  nil,
		}
		return res, nil
	}

	successMsg := "成功获取视频"
	res := &favorite.FavoriteListResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  &successMsg,
		VideoList:  videoListProto,
	}
	return res, nil
}
