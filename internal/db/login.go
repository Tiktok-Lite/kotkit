package db

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func CreateLogin(login *model.Login) error {
	if err := DB().Create(login).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}

func UpdateLogin(login *model.Login) error {
	if err := DB().Save(login).Error; err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func QueryLoginByName(name string) (*model.Login, error) {
	var login model.Login
	if err := DB().Where("Username = ?", name).First(&login).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("failed to query user by name")
	}
	return &login, nil
}
