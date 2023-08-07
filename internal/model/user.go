package model

type User struct {
	ID              int64 `gorm:"primaryKey"`
	Name            string
	FollowCount     uint
	IsFollow        bool
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  uint
	WorkCount       uint
	FavoriteCount   uint
	UserLogin       Login     // user和login一对一
	Comments        []Comment // 用户和评论一对多
	Followers       []User    `gorm:"many2many:user_relations;"`   // 用户之间多对多，使用中间表user_relations
	LikedVideos     []Video   `gorm:"many2many:user_like_videos;"` // 用户和视频多对多，使用中间表user_like_videos
}
