package rpc

import (
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
)

func InitRPC() {
	userConfig := conf.LoadConfig(constant.DefaultUserConfigName)
	InitUser(userConfig)
}
