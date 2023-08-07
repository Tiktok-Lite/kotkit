package rpc

import (
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/viper"
)

func InitRPC() {
	userConfig := viper.LoadConfig(constant.DefaultUserConfigPath)
	InitUser(userConfig)
}
