package response

import "github.com/Tiktok-Lite/kotkit/kitex_gen/video"

type Feed struct {
	Base
	NextTime  *int64         `json:"next_time"`
	VideoList []*video.Video `json:"video_list"`
}

type PublishList struct {
	Base
	VideoList []*video.Video `json:"video_list"`
}

func PackFeedError(errorMsg string) Feed {
	base := PackBaseError(errorMsg)
	return Feed{
		Base:      base,
		NextTime:  nil,
		VideoList: nil,
	}
}

func PackFeedSuccess(nextTime *int64, videoList []*video.Video, msg string) Feed {
	base := PackBaseSuccess(msg)
	return Feed{
		Base:      base,
		NextTime:  nextTime,
		VideoList: videoList,
	}
}

func PackPublishListError(errorMsg string) PublishList {
	base := PackBaseError(errorMsg)
	return PublishList{
		Base:      base,
		VideoList: nil,
	}
}

func PackPublishListSuccess(videoList []*video.Video, msg string) PublishList {
	base := PackBaseSuccess(msg)
	return PublishList{
		Base:      base,
		VideoList: videoList,
	}
}
