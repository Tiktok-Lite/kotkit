package response

import (
	"github.com/Tiktok-Lite/kotkit/kitex_gen/user"
)

type RelationList struct {
	Base
	UserList []*user.User `json:"user_list"`
}

type Relation struct {
	Base
}

func PackListSuccess(userList []*user.User, msg string) RelationList {
	base := PackBaseSuccess(msg)
	return RelationList{
		Base:     base,
		UserList: userList,
	}
}

func PackListError(errorMsg string) RelationList {
	base := PackBaseError(errorMsg)
	return RelationList{
		Base:     base,
		UserList: nil,
	}
}

func PackActionSuccess(msg string) Relation {
	base := PackBaseSuccess(msg)
	return Relation{
		Base: base,
	}
}
func PackActionError(errorMsg string) Relation {
	base := PackBaseError(errorMsg)
	return Relation{
		Base: base,
	}
}
