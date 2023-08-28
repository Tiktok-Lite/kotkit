package response

import "github.com/Tiktok-Lite/kotkit/kitex_gen/video"

type FavoriteAction struct {
	Base
}

type FavoriteList struct {
	Base
	VideoList []*video.Video `json:"video_list"`
}

func PackFavoriteActionError(errorMsg string) FavoriteAction {
	base := PackBaseError(errorMsg)
	return FavoriteAction{
		Base: base,
	}
}

func PackFavoriteActionSuccess(msg string) FavoriteAction {
	base := PackBaseSuccess(msg)
	return FavoriteAction{
		Base: base,
	}
}

func PackFavoriteListError(errorMsg string) FavoriteList {
	base := PackBaseError(errorMsg)
	return FavoriteList{
		Base:      base,
		VideoList: nil,
	}
}

func PackFavoriteListSuccess(videoList []*video.Video, msg string) FavoriteList {
	base := PackBaseSuccess(msg)
	return FavoriteList{
		Base:      base,
		VideoList: videoList,
	}
}
