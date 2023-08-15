package response

// 定义一个基用来返回http中的状态码和消息
// kitex 生成的json不好用，会有bug，这里再封装一下

type Base struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
