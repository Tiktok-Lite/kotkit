package repository

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type LoginRepository interface {
	CreateLogin(login *model.Login) error
	UpdateLogin(login *model.Login) error
	QueryLoginByName(name string) (*model.Login, error)
}

type loginRepository struct {
	*Repository
}

func NewLoginRepository(r *Repository) LoginRepository {
	return &loginRepository{
		Repository: r,
	}
}

func (r *loginRepository) CreateLogin(login *model.Login) error {
	// TODO(century): add error info to log
	if err := r.db.Create(login).Error; err != nil {
		return errors.Wrap(err,"failed to create user")
	}

	return nil
}

func (r *loginRepository) UpdateLogin(login *model.Login) error {
	if err := r.db.Save(login).Error; err != nil {
		return errors.Wrap(err,"failed to update user")
	}

	return nil
}


func (r *loginRepository) QueryLoginByName(name string) (*model.Login, error) {
    var login model.Login
    if err := r.db.Where("Username = ?", name).First(&login).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
		return nil, errors.New("failed to query user by name")
    }
    return &login, nil
}

