package model

import "gorm.io/gorm"

type Relation struct {
	gorm.DB
	UserId   uint
	ToUserId uint
}
