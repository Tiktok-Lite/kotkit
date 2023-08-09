package repository

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/pkg/errors"
)

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	QueryUserByID(id int64) (*model.User, error)
	QueryUserByName(name string) (*model.User, error)
}

type userRepository struct {
	*Repository
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

func (r *userRepository) Create(user *model.User) error {
	// TODO(century): add error info to log
	if err := r.db.Create(user).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (r *userRepository) Update(user *model.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func (r *userRepository) QueryUserByID(id int64) (*model.User, error) {
	var user model.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("failed to query user")
	}

	return &user, nil
}

func (r *userRepository) QueryUserByName(name string) (*model.User, error) {
    var user model.User
    if err := r.db.Where("name = ?", name).First(&user).Error; err != nil {
        return nil, errors.New("failed to query user by name")
    }
    return &user, nil
}


