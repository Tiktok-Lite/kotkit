package main

import (
	"context"
	"fmt"
	"github.com/Tiktok-Lite/kotkit/cmd/comment/pack"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/comment"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/converter"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/Tiktok-Lite/kotkit/pkg/oss"
	"time"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	logger := log.Logger()

	claims, err := Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("Error occurs when parsing token. %v", err)
		res := &comment.DouyinCommentActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	if req.ActionType != constant.PostCommentCode && req.ActionType != constant.DeleteCommentCode {
		res := &comment.DouyinCommentActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "ActionType错误",
			Comment:    nil,
		}
		return res, nil
	}

	// 以mm-dd的格式获取当前时间, 类型为字符串
	date := time.Now()
	c := model.Comment{
		Content:    req.CommentText,
		VideoID:    uint(req.VideoId),
		UserID:     uint(claims.Id),
		CreateDate: fmt.Sprintf("%2d-%2d", date.Month(), date.Day()),
	}
	if req.ActionType == constant.PostCommentCode {
		err := db.CommentTransaction(&c)
		if err != nil {
			logger.Errorf("Error occurs when add comment to database. %v", err)
			res := &comment.DouyinCommentActionResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "发布评论失败",
				Comment:    nil,
			}
			return res, nil
		}
	}
	if req.ActionType == constant.DeleteCommentCode {
		err := db.UnCommentTransaction(req.CommentId, c.VideoID)
		if err != nil {
			logger.Errorf("Error occurs when delete comment to database. %v", err)
			res := &comment.DouyinCommentActionResponse{
				StatusCode: constant.StatusErrorCode,
				StatusMsg:  "删除评论失败",
				Comment:    nil,
			}
			return res, nil
		}
	}
	usr, err := db.QueryUserByID(int64(c.UserID))
	if err != nil {
		logger.Errorf("Error occurs when querying user from database. %v", err)
		res := &comment.DouyinCommentActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "查询用户失败",
			Comment:    nil,
		}
		return res, nil
	}

	avatarUrl, err := oss.GetObjectURL(oss.AvatarBucketName, usr.Avatar)
	if err != nil {
		logger.Errorf("Failed to get object url due to %v", err)
		res := &comment.DouyinCommentActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "minio数据库查询错误",
			Comment:    nil,
		}
		return res, nil
	}
	usr.Avatar = avatarUrl

	bgImgUrl, err := oss.GetObjectURL(oss.BackgroundImageBucketName, usr.BackgroundImage)
	if err != nil {
		logger.Errorf("Failed to get object url due to %v", err)
		res := &comment.DouyinCommentActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "minio数据库查询错误",
			Comment:    nil,
		}
		return res, nil
	}
	usr.BackgroundImage = bgImgUrl

	cmtProto, err := converter.ConvertCommentModelToProto(&c, usr)
	if err != nil {
		logger.Errorf("Error occurs when converting comment model to proto. %v", err)
		res := &comment.DouyinCommentActionResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "评论操作失败",
			Comment:    nil,
		}
		return res, nil
	}

	res := &comment.DouyinCommentActionResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "评论操作成功",
		Comment:    cmtProto,
	}
	return res, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	_, err = Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("Error occurs when parsing token. %v", err)
		res := &comment.DouyinCommentListResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	res := &comment.DouyinCommentListResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "成功获取评论列表",
	}
	comments, err := db.QueryCommentByVideoID(req.VideoId)
	if err != nil {
		logger.Errorf("Error occurs when querying comment list from database. %v", err)
		res.StatusCode = constant.StatusErrorCode
		res.StatusMsg = "查询评论列表失败"
		return res, nil
	}
	commentList := pack.CommentList(comments)
	res.CommentList = commentList

	return res, nil
}
