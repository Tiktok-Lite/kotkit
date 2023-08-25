package repository

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RelationRepository interface {
	NewRelation(relation *model.Relation) error
	DelRelation(relation *model.Relation) error
	QueryRelationByID(userID uint, toUserID uint) (*model.Relation, error)
	GetFollowerListByUserID(toUserID uint) ([]*model.Relation, error)
	GetFollowingListByUserID(UserID uint) ([]*model.Relation, error)
	GetFriendList(toUserID uint) ([]*model.Relation, error)
}

type relationRepository struct {
	*Repository
}

func NewRelationRepository(r *Repository) RelationRepository {
	return &relationRepository{
		Repository: r,
	}
}

func (r *relationRepository) NewRelation(relation *model.Relation) error {
	if err := r.db.Create(relation).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *relationRepository) DelRelation(relation *model.Relation) error {
	if err := r.db.Delete(relation).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}

func (r *relationRepository) QueryRelationByID(userID uint, toUserID uint) (*model.Relation, error) {
	var relation model.Relation
	if err := r.db.Where("user_id = ? AND to_user_id IN ?", userID, toUserID).First(&relation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("failed to query user by id")
	}
	return &relation, nil
}

func (r *relationRepository) GetFollowerListByUserID(toUserID uint) ([]*model.Relation, error) {
	var RelationList []*model.Relation
	if err := r.db.Where("to_user_id = ?", toUserID).Find(&RelationList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("failed to query user by id")
	}
	return RelationList, nil
}

func (r *relationRepository) GetFollowingListByUserID(UserID uint) ([]*model.Relation, error) {
	var RelationList []*model.Relation
	if err := r.db.Where("user_id = ?", UserID).Find(&RelationList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("failed to query user by id")
	}
	return RelationList, nil
}

func (r *relationRepository) GetFriendList(UserID uint) ([]*model.Relation, error) {
	var RelationList []*model.Relation
	if err := r.db.Where("user_id = ?", UserID).Find(&RelationList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("failed to query user by id")
	}
	return RelationList, nil
}
