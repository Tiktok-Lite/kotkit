package response

import "github.com/Tiktok-Lite/kotkit/kitex_gen/user"

type UserInfo struct {
	Base
	User *user.User `json:"user"`
}

func PackUserInfoSuccess(userInfo *user.User, msg string) UserInfo {
	base := PackBaseSuccess(msg)
	return UserInfo{
		Base: base,
		User: userInfo,
	}
}

func PackUserInfoError(errorMsg string) UserInfo {
	base := PackBaseError(errorMsg)
	return UserInfo{
		Base: base,
		User: nil,
	}
}
