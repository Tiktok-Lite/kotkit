package model

type User struct {
	ID              uint   `gorm:"id, primary_key"`
	Name            string `gorm:"name"`
	FollowCount     uint   `gorm:"follow_count"`
	IsFollow        bool   `gorm:"is_follow"`
	Avatar          string `gorm:"avatar"`
	BackgroundImage string `gorm:"background_image"`
	Signature       string `gorm:"signature"`
	TotalFavorited  uint   `gorm:"total_favorited"`
	WordCount       uint   `gorm:"word_count"`
	FavoriteCount   uint   `gorm:"favorite_count"`
}
