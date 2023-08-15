package response

import (
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
)

// 定义一个基用来返回http中的状态码和消息
// kitex 生成的json不好用，会有bug，这里再封装一下

type Base struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func PackBaseSuccess(msg string) Base {
	return Base{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  msg,
	}
}

func PackBaseError(msg string) Base {
	return Base{
		StatusCode: constant.StatusErrorCode,
		StatusMsg:  msg,
	}
}
