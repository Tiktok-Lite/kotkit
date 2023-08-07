package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	User       User
	Content    string
	CreateDate string
	UserID     uint // 用户与评论之间的多对一
	VideoID    uint // 评论和视频之间多对一
}
