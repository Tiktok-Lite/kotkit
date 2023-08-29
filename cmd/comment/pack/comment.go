package pack

import (
	"errors"
	"github.com/Tiktok-Lite/kotkit/internal/db"
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/comment"
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user"
	"gorm.io/gorm"
)

func User(u *model.User) *user.User {
	return &user.User{
		Id:              int64(u.ID),
		Name:            u.Name,
		FollowCount:     &u.FollowCount,
		FollowerCount:   &u.FollowerCount,
		IsFollow:        u.IsFollow,
		Avatar:          &u.Avatar,
		BackgroundImage: &u.BackgroundImage,
		Signature:       &u.Signature,
		TotalFavorited:  &u.TotalFavorited,
		WorkCount:       &u.WorkCount,
		FavoriteCount:   &u.FavoriteCount,
	}
}

// Note pack note info
func Comment(m *model.Comment) *comment.Comment {
	if m == nil {
		return nil
	}
	usr, err := db.QueryUserByID(int64(m.UserID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	u := User(usr)
	return &comment.Comment{
		Id:         int64(m.ID),
		User:       u,
		Content:    m.Content,
		CreateDate: m.CreatedAt.Format("01-02"),
	}
}
func CommentList(cs []*model.Comment) []*comment.Comment {
	comments := make([]*comment.Comment, 0)
	for _, c := range cs {
		if n := Comment(c); n != nil {
			comments = append(comments, n)
		}
	}
	return comments
}
