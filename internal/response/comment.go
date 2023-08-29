package response

import "github.com/Tiktok-Lite/kotkit/kitex_gen/comment"

type CommentList struct {
	Base
	CommentList []*comment.Comment `json:"comment"`
}

type Comment struct {
	Base
	Comment *comment.Comment `json:"comment"`
}

func PackCommentListSuccess(commentList []*comment.Comment, msg string) CommentList {
	base := PackBaseSuccess(msg)
	return CommentList{
		Base:        base,
		CommentList: commentList,
	}
}

func PackCommentActionSuccess(comment *comment.Comment, msg string) Comment {
	base := PackBaseSuccess(msg)
	return Comment{
		Base:    base,
		Comment: comment,
	}
}
