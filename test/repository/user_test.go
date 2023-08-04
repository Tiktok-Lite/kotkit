package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func setupRepository(t *testing.T) (repository.UserRepository, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := repository.NewRepository(db)
	userRepo := repository.NewUserRepository(repo)

	return userRepo, mock
}

func TestUserRepositoryCreate(t *testing.T) {
	userRepo, mock := setupRepository(t)

	user := &model.User{
		Name:            "test",
		FollowCount:     1,
		IsFollow:        false,
		Avatar:          "avatar",
		BackgroundImage: "background_image",
		Signature:       "signature",
		TotalFavorited:  1,
		WordCount:       1,
		FavoriteCount:   1,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").
		WithArgs(user.Name, user.FollowCount, user.IsFollow, user.Avatar, user.BackgroundImage,
			user.Signature, user.TotalFavorited, user.WordCount, user.FavoriteCount).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := userRepo.Create(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepositoryUpdate(t *testing.T) {
	userRepo, mock := setupRepository(t)

	user := &model.User{
		ID:              1,
		Name:            "test",
		FollowCount:     1,
		IsFollow:        false,
		Avatar:          "avatar",
		BackgroundImage: "background_image",
		Signature:       "signature",
		TotalFavorited:  1,
		WordCount:       1,
		FavoriteCount:   1,
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users`").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := userRepo.Update(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepositoryQueryUserByID(t *testing.T) {
	userRepo, mock := setupRepository(t)

	user := &model.User{
		ID:              1,
		Name:            "test",
		FollowCount:     1,
		IsFollow:        false,
		Avatar:          "avatar",
		BackgroundImage: "background_image",
		Signature:       "signature",
		TotalFavorited:  1,
		WordCount:       1,
		FavoriteCount:   1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "follow_count", "is_follow", "avatar", "background_image",
		"signature", "total_favorited", "word_count", "favorite_count"}).
		AddRow(user.ID, user.Name, user.FollowCount, user.IsFollow, user.Avatar, user.BackgroundImage,
			user.Signature, user.TotalFavorited, user.WordCount, user.FavoriteCount)
	mock.ExpectQuery("SELECT").
		WithArgs(user.ID).
		WillReturnRows(rows)

	result, err := userRepo.QueryUserByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}
