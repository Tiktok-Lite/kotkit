package response

import "github.com/Tiktok-Lite/kotkit/kitex_gen/user"

type UserInfo struct {
	Base
	User *user.User `json:"user"`
}
