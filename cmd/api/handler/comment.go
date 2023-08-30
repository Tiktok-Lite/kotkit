package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/comment"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
	"strings"
)

func CommentList(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	videoId := c.Query("video_id")
	if videoId == "" {
		logger.Errorf("Illegal input: empty video_id.")
		ResponseError(c, http.StatusBadRequest, response.PackCommentListError("video_id不能为空"))
		return
	}
	id, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse video_id: %v", err)
		ResponseError(c, http.StatusBadRequest, response.PackCommentActionError("请检查您的输入是否合法"))
		return
	}

	token := c.Query("token")
	if token == "" {
		logger.Errorf("Illegal input: empty token.")
		ResponseError(c, http.StatusBadRequest, response.PackCommentListError("token不能为空"))
		return
	}
	req := &comment.DouyinCommentListRequest{
		Token:   token,
		VideoId: id,
	}
	resp, err := rpc.CommentList(ctx, req)
	if err != nil {
		logger.Errorf("error occurs when calling rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError, response.PackCommentListError(resp.StatusMsg))
		return
	}
	ResponseSuccess(c, response.PackCommentListSuccess(resp.CommentList, "评论列表获取成功"))
}

func CommentAction(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	videoId := c.Query("video_id")
	if videoId == "" {
		logger.Errorf("Illegal input: empty video_id.")
		ResponseError(c, http.StatusBadRequest, response.PackCommentActionError("video_id不能为空"))
		return
	}
	id, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse video_id: %v", err)
		ResponseError(c, http.StatusBadRequest, response.PackCommentActionError("请检查video_id是否合法"))
		return
	}

	token := c.Query("token")
	if token == "" {
		logger.Errorf("Illegal input: empty token.")
		ResponseError(c, http.StatusBadRequest, response.PackCommentActionError("token不能为空"))
		return
	}

	actionType := c.Query("action_type")
	if actionType == "" {
		logger.Errorf("Illegal input: empty action_type.")
		ResponseError(c, http.StatusBadRequest, response.PackCommentActionError("action_type不能为空"))
		return
	}
	act, err := strconv.Atoi(actionType)
	if err != nil {
		logger.Errorf("failed to parse action_type: %v", err)
		ResponseError(c, http.StatusBadRequest, response.PackCommentActionError("请检查您的输入是否合法"))
		return
	}

	req := &comment.DouyinCommentActionRequest{
		Token:      token,
		VideoId:    id,
		ActionType: int32(act),
	}
	if act == constant.PostCommentCode {
		content := c.Query("comment_text")
		s := strings.TrimSpace(content)
		if len(s) == 0 {
			ResponseError(c, http.StatusInternalServerError, response.PackCommentActionError("评论不能为空"))
			return
		}
		req.CommentText = s
	}
	if act == constant.DeleteCommentCode {
		commentId := c.Query("comment_id")
		if commentId == "" {
			logger.Errorf("Illegal input: empty comment_id.")
			ResponseError(c, http.StatusBadRequest, response.PackCommentActionError("comment_id不能为空"))
			return
		}
		cid, err := strconv.ParseInt(commentId, 10, 64)
		if err != nil {
			ResponseError(c, http.StatusInternalServerError, response.PackCommentActionError("请检查comment_id是否合法"))
			return
		}
		req.CommentId = cid
	}

	resp, err := rpc.CommentAction(ctx, req)
	if err != nil {
		logger.Errorf("failed to call rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError,
			response.PackCommentActionError(resp.StatusMsg))
		return
	}

	ResponseSuccess(c, response.PackCommentActionSuccess(resp.Comment, resp.StatusMsg))
}
