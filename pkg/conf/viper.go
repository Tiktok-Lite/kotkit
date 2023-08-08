package conf

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
}

func LoadConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	conf.AddConfigPath("./")
	conf.AddConfigPath("../")
	conf.AddConfigPath("../../")

	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
