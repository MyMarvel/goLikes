package settings

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.AddConfigPath("./settings")
	viper.SetConfigName("default")

	return viper.ReadInConfig()
}
