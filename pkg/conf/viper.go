package conf

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
}

func LoadConfig(configName string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigType("yml")
	conf.SetConfigName(configName)
	// 使用脚本启动时会进入三层路径，故这样设置避免出现问题，不是一个优雅的处理方法
	conf.AddConfigPath("./config")
	conf.AddConfigPath("../config")
	conf.AddConfigPath("../../config")
	conf.AddConfigPath("../../../config")

	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
