package rpc

import (
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
)

var (
	logger = log.Logger()
)

func InitRPC() {
	userConfig := conf.LoadConfig(constant.DefaultUserConfigName)
	InitUser(userConfig)

	videoConfig := conf.LoadConfig(constant.DefaultVideoConfigName)
	InitVideo(videoConfig)

	loginConfig := conf.LoadConfig(constant.DefaultLoginConfigName)
	InitLogin(loginConfig)

	favoriteConfig := conf.LoadConfig(constant.DefaultFavoriteConfigName)
	InitFavorite(favoriteConfig)

	relationConfig := conf.LoadConfig(constant.DefaultRelationConfigName)
	InitRelation(relationConfig)
}
