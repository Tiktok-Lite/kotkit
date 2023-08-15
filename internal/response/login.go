package response

type Login struct {
	Base
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func PackLoginOrRegisterSuccess(userID int64, token string, msg string) Login {
	base := PackBaseSuccess(msg)
	return Login{
		Base:   base,
		UserID: userID,
		Token:  token,
	}
}

func PackLoginOrRegisterError(errorMsg string) Login {
	base := PackBaseError(errorMsg)
	return Login{
		Base:   base,
		UserID: -1,
		Token:  "",
	}
}
