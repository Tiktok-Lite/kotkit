package command

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// RelationAction action favorite.
func (*RelationActionService) RelationAction(UserId int64, req *relation.RelationActionRequest) error {
	// 1-关注
	if req.ActionType == 1 {
		return db.NewRelation(uint(UserId), uint(req.ToUserId))
	}
	// 2-取消关注
	if req.ActionType == 2 {
		return db.DelRelation(uint(UserId), uint(req.ToUserId))
	}
	return nil
}
