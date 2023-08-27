package db

import (
	"errors"
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"gorm.io/gorm"
)

func Create(user *model.User) error {
	// TODO(century): add error info to log
	if err := DB().Create(user).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func Update(user *model.User) error {
	if err := DB().Save(user).Error; err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func UpdateByUsername(username string, updatedUser *model.User) error {
	// 构建更新条件
	condition := map[string]interface{}{
		"Name": username,
	}
	// 执行更新操作
	if err := DB().Model(&model.User{}).Where(condition).Updates(updatedUser).Error; err != nil {
		return errors.New("failed to update user by username")
	}

	return nil
}

func QueryUserByID(id int64) (*model.User, error) {
	var user model.User
	if err := DB().Where("id = ?", id).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func QueryUserByName(name string) (*model.User, error) {
	var user model.User
	if err := DB().Where("name = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("failed to query user by name")
	}
	return &user, nil
}

func QueryUserByRelation(userID, followerID int64) (bool, error) {
	var count int64
	err := DB().Raw("SELECT COUNT(*) FROM user_relations WHERE user_id = ? AND follower_id = ?", userID, followerID).Count(&count).Error
	if err != nil {
		return false, errors.New("failed to query user by relation")
	}
	return count > 0, nil
}
