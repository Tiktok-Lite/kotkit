package conf

import (
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/spf13/viper"
)

func NewConfig() *viper.Viper {
	return getConfig(constant.DefaultConfigPath)
}

// parse file
func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
