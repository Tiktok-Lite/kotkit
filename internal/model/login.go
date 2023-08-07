package model

import "gorm.io/gorm"

type Login struct {
	gorm.Model
	Username string
	Password string `gorm:"not null"`
	UserID   uint   // login和user之间一对一
}
