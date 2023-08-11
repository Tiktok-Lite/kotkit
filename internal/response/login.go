package response


type Login struct {
	Base
	UserID	int64         `json:"user_id"`
	Token	string 		  `json:"token"`
}
