package response

import (
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user"
)

type Relation struct {
	Base
	UserList []*user.User `json:"user_list"`
}

func PackListSuccess(userList []*user.User, msg string) Relation {
	base := PackBaseSuccess(msg)
	return Relation{
		Base:     base,
		UserList: userList,
	}
}

func PackListError(errorMsg string) Relation {
	base := PackBaseError(errorMsg)
	return Relation{
		Base:     base,
		UserList: nil,
	}
}
