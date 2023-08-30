package main

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/comment/pack"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/comment"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
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
	res := &comment.DouyinCommentActionResponse{
		StatusCode: constant.StatusErrorCode,
		StatusMsg:  "success",
		Comment:    nil,
	}
	if req.ActionType == constant.PostCommentCode {
		c := model.Comment{
			Content: req.CommentText,
			VideoID: uint(req.VideoId),
			UserID:  uint(claims.Id),
		}
		err := db.AddComment(&c)
		if err != nil {
			logger.Errorf("Error occurs when add comment to database. %v", err)
			res.StatusCode = constant.StatusErrorCode
			res.StatusMsg = "发布评论失败"
		}
		return res, err
	}
	if req.ActionType == constant.DeleteCommentCode {
		err := db.DeleteCommentById(req.CommentId)
		if err != nil {
			logger.Errorf("Error occurs when delete comment to database. %v", err)
			res.StatusCode = constant.StatusErrorCode
			res.StatusMsg = "删除评论失败"
		}
		return res, nil
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
