package converter

import (
	"github.com/Tiktok-Lite/kotkit/internal/model"
	genUser "github.com/Tiktok-Lite/kotkit/kitex_gen/user"
	"github.com/pkg/errors"
)

// TODO: 将model转换为proto生成后的类型，以简化版handler中的处理
func ConvertVideoModelListToProto(videoList []*model.Video) {

}

func ConvertUserModelToProto(user *model.User) (*genUser.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}
	return &genUser.User{
		Id:              user.ID,
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
