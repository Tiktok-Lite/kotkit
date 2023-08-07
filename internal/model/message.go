package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ToUserID   uint
	FromUserID uint
	Content    string
	CreateTime string
}
