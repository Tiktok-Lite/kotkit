package db

import (
	"fmt"
	"github.com/Tiktok-Lite/kotkit/internal/model"
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
			return errors.Wrap(err, "failed to create follow relation")
		}

		// 2. 改变 user 表中的 following count
		if err := tx.Model(&model.User{}).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return errors.Wrap(err, "failed to update following count")
		}

		// 3. 改变 user 表中的 follower count
		if err := tx.Model(&model.User{}).Where("id = ?", toUserID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
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
			return errors.Wrap(err, "failed to create follow relation")
		}
		// 2. 改变 user 表中的 following count
		if err := tx.Model(&model.User{}).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			return errors.Wrap(err, "failed to update following count")
		}

		// 3. 改变 user 表中的 follower count
		if err := tx.Model(&model.User{}).Where("id = ?", toUserID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return errors.Wrap(err, "failed to update follower count")
		}

		return nil
	})
	return err
}

func QueryRelationByID(userID uint, followerID uint) (*model.User, error) {
	var relation model.User
	if err := DB().Where("user_id = ? AND follower_id IN ?", userID, followerID).First(&relation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("failed to query user by id")
	}
	return &relation, nil
}

func GetFollowerListByUserID(UserID uint) ([]*FollowRelation, error) {
	var RelationList []*FollowRelation
	fmt.Println(UserID)
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
