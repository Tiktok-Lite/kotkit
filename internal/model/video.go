package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Author        User `gorm:"-"`
	PlayURL       string
	CoverURL      string
	FavoriteCount uint
	CommentCount  uint
	IsFavorite    bool
	Title         string
	UserID        uint // 表示视频和用户一对多关系
	Comments      []Comment
	Users         []User `gorm:"many2many:user_like_videos;"`
}
