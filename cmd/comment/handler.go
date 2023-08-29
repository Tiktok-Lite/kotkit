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

	res := &comment.DouyinCommentActionResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "success",
		Comment:    nil,
	}
	claims, err := Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("Error occurs when parsing token. %v", err)
		res.StatusMsg = "token错误"
		return res, err
	}
	if req.ActionType != 1 && req.ActionType != 2 {
		res.StatusCode = constant.StatusErrorCode
		res.StatusMsg = "ActionType错误"
		return res, nil
	}
	if req.ActionType == 1 {
		c := model.Comment{
			Content: req.CommentText,
			VideoID: uint(req.VideoId),
			UserID:  uint(claims.Id),
		}
		err := db.AddComment(&c)
		if err != nil {
			logger.Errorf("Error occurs when add comment to database. %v", err)
			res.StatusMsg = "评论添加失败"
		}

		return res, err
	}
	if req.ActionType == 2 {
		err := db.DeleteCommentById(req.CommentId)
		if err != nil {
			logger.Errorf("Error occurs when delete comment to database. %v", err)
			res.StatusMsg = "评论删除失败"
		}
		return res, err
	}

	return res, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	res := &comment.DouyinCommentListResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "成功获取评论",
	}

	_, err = Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("Error occurs when parsing token. %v", err)
		res.StatusMsg = "token错误"
		return res, err
	}

	comments, err := db.QueryCommentByVideoID(req.VideoId)

	if err != nil {
		logger.Errorf("Error occurs when querying comment list from database. %v", err)
		return nil, err
	}
	commentList := pack.CommentList(comments)
	res.CommentList = commentList

	return res, nil
}
