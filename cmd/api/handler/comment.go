package handler

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/cmd/api/rpc"
	"github.com/Tiktok-Lite/kotkit/internal/response"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/comment"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
	"strings"
)

func CommentList(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()

	videoID := c.Query("video_id")
	token := c.Query("token")
	id, err := strconv.ParseInt(videoID, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse video_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackBaseError("请检查您的输入是否合法"))
		return
	}
	req := &comment.DouyinCommentListRequest{
		Token:   token,
		VideoId: id,
	}
	resp, err := rpc.CommentList(ctx, req)
	if err != nil {
		logger.Errorf("failed to call rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError,
			response.PackBaseError("评论信息获取失败，服务器内部问题"))
		return
	}

	ResponseSuccess(c, response.PackCommentListSuccess(resp.CommentList, "评论信息获取成功"))
}

func CommentAction(ctx context.Context, c *app.RequestContext) {
	logger := log.Logger()
	token := c.Query("token")

	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	id, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		logger.Errorf("failed to parse video_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackBaseError("请检查您的输入是否合法"))
		return
	}
	act, err := strconv.Atoi(actionType)
	if err != nil {
		logger.Errorf("failed to parse video_id: %v", err)
		ResponseError(c, http.StatusBadRequest,
			response.PackBaseError("请检查您的输入是否合法"))
		return
	}

	req := &comment.DouyinCommentActionRequest{
		Token:      token,
		VideoId:    id,
		ActionType: int32(act),
	}
	msg := ""
	if act == 1 {
		content := c.Query("comment_text")
		s := strings.TrimSpace(content)
		if len(s) == 0 {
			ResponseError(c, http.StatusInternalServerError,
				response.PackBaseError("获取失败"))
			return
		}
		req.CommentText = s
		msg = "评论成功"
	} else {
		commentId := c.Query("comment_id")
		cid, err := strconv.Atoi(commentId)
		if err != nil {
			ResponseError(c, http.StatusInternalServerError,
				response.PackBaseError("获取失败"))
			return
		}
		cid64 := int64(cid)
		req.CommentId = cid64
		msg = "删除评论成功"
	}

	resp, err := rpc.CommentAction(ctx, req)
	if err != nil {
		logger.Errorf("failed to call rpc: %v", err)
		ResponseError(c, http.StatusInternalServerError,
			response.PackBaseError("评论操作失败，服务器内部问题"))
		return
	}

	ResponseSuccess(c, response.PackCommentActionSuccess(resp.Comment, msg))
}
