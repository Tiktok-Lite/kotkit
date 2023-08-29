package db

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/relation"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FollowRelation struct {
	UserID     uint `gorm:"index:idx_userid;not null"`
	FollowerID uint `gorm:"index:idx_userid;index:idx_userid_to;not null"`
}

func (FollowRelation) TableName() string {
	return "user_relations"
}

func NewRelation(userID uint, toUserID uint) error {
	err := DB().Transaction(func(tx *gorm.DB) error {
		// 1. 新增关注数据
		err := tx.Create(&FollowRelation{FollowerID: userID, UserID: toUserID}).Error
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "failed to create follow relation")
		}

		// 2. 改变 user 表中的 following count
		if err := tx.Model(&model.User{}).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			tx.Rollback()
			return errors.Wrap(err, "failed to update following count")
		}

		// 3. 改变 user 表中的 follower count
		if err := tx.Model(&model.User{}).Where("id = ?", toUserID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			tx.Rollback()
			return errors.Wrap(err, "failed to update follower count")
		}

		return nil
	})
	return err
}

func DelRelation(userID uint, toUserID uint) error {
	err := DB().Transaction(func(tx *gorm.DB) error {
		relation := new(FollowRelation)
		// 1. 删除关注数据
		err := tx.Unscoped().Where("user_id = ? AND follower_id = ?", toUserID, userID).Delete(&relation).Error
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "failed to create follow relation")
		}
		// 2. 改变 user 表中的 following count
		if err := tx.Model(&model.User{}).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			tx.Rollback()
			return errors.Wrap(err, "failed to update following count")
		}

		// 3. 改变 user 表中的 follower count
		if err := tx.Model(&model.User{}).Where("id = ?", toUserID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			tx.Rollback()
			return errors.Wrap(err, "failed to update follower count")
		}

		return nil
	})
	return err
}

func GetFollowerListByUserID(UserID uint) ([]*FollowRelation, error) {
	var RelationList []*FollowRelation
	if err := DB().Where("user_id = ?", UserID).Find(&RelationList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("failed to query user by id")
	}
	return RelationList, nil
}

func GetFollowingListByUserID(UserID uint) ([]*FollowRelation, error) {
	var RelationList []*FollowRelation
	if err := DB().Where("follower_id = ?", UserID).Find(&RelationList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("failed to query user by id")
	}
	return RelationList, nil
}

func GetFriendList(UserID uint) ([]*FollowRelation, error) {
	var FriendList []*FollowRelation
	if err := DB().Raw("SELECT user_id, follower_id FROM user_relations WHERE user_id = ? AND follower_id IN (SELECT user_id FROM user_relations r WHERE r.follower_id = user_relations.user_id)", UserID).Scan(&FriendList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("failed to query user by id")
	}
	return FriendList, nil
}

func RelationAction(UserId int64, req *relation.RelationActionRequest) error {
	// 1-关注
	if req.ActionType == 1 {
		return NewRelation(uint(UserId), uint(req.ToUserId))
	}
	// 2-取消关注
	if req.ActionType == 2 {
		return DelRelation(uint(UserId), uint(req.ToUserId))
	}
	return nil
}
