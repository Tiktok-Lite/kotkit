package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string
	FollowCount     int64
	FollowerCount   int64
	IsFollow        bool
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  int64
	WorkCount       int64
	FavoriteCount   int64
	UserLogin       Login     // user和login一对一
	Comments        []Comment // 用户和评论一对多
	Followers       []User    `gorm:"many2many:user_relations;"`   // 用户之间多对多，使用中间表user_relations
	LikedVideos     []Video   `gorm:"many2many:user_like_videos;"` // 用户和视频多对多，使用中间表user_like_videos
}
