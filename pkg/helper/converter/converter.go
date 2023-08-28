package converter

import (
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/internal/model"
	genUser "github.com/Tiktok-Lite/kotkit/kitex_gen/user"
	genVideo "github.com/Tiktok-Lite/kotkit/kitex_gen/video"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func ConvertVideoModelListToProto(videoList []*model.Video) ([]*genVideo.Video, error) {
	if videoList == nil {
		return nil, errors.New("video list is nil")
	}
	var res []*genVideo.Video
	for _, video := range videoList {
		userProto, _ := ConvertUserModelToProto(&video.Author)
		res = append(res, &genVideo.Video{
			Id:            int64(video.ID),
			Author:        userProto,
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}

	return res, nil
}

func ConvertFollowingListModelToProto(userList []*db.FollowRelation) ([]*genUser.User, error) {
	users := make([]*genUser.User, 0)
	if userList == nil {
		return nil, errors.New("user is nil")
	}
	for _, user := range userList {
		user2, err := db.QueryUserByID(int64(user.UserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		userProto, _ := ConvertUserModelToProto(user2)
		users = append(users, userProto)
	}
	return users, nil
}

func ConvertFollowerListModelToProto(userList []*db.FollowRelation) ([]*genUser.User, error) {
	users := make([]*genUser.User, 0)
	if userList == nil {
		return nil, errors.New("user is nil")
	}
	for _, user := range userList {
		user2, err := db.QueryUserByID(int64(user.FollowerID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		userProto, _ := ConvertUserModelToProto(user2)
		users = append(users, userProto)
	}
	return users, nil
}

func ConvertUserModelToProto(user *model.User) (*genUser.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}
	return &genUser.User{
		Id:              int64(user.ID),
		Name:            user.Name,
		FollowCount:     &user.FollowCount,
		FollowerCount:   &user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          &user.Avatar,
		BackgroundImage: &user.BackgroundImage,
		Signature:       &user.Signature,
		TotalFavorited:  &user.TotalFavorited,
		WorkCount:       &user.WorkCount,
		FavoriteCount:   &user.FavoriteCount,
	}, nil
}
