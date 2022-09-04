package configs

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("default")

	return viper.ReadInConfig()
}
