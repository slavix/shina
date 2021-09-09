package config

import (
	"github.com/spf13/viper"
	"github.com/ztrue/tracerr"
)

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return tracerr.Wrap(viper.ReadInConfig())
}

func GetString(name string) string {
	return viper.GetString(name)
}
