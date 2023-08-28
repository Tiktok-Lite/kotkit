package command

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation"
)

type RelationActionService struct {
	ctx context.Context
}

var (
	repo         = repository.NewRepository(db.DB())
	relationRepo = repository.NewRelationRepository(repo)
)

// NewRelationActionService new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// RelationAction action favorite.
func (*RelationActionService) RelationAction(UserId int64, req *relation.RelationActionRequest) error {
	// 1-关注
	if req.ActionType == 1 {
		return relationRepo.NewRelation(uint(UserId), uint(req.ToUserId))
	}
	// 2-取消关注
	if req.ActionType == 2 {
		return relationRepo.DelRelation(uint(UserId), uint(req.ToUserId))
	}
	return nil
}
