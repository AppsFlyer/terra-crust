package app

import (
	"os"

	logger "github.com/AppsFlyer/go-logger"
	"github.com/spf13/viper"
)

func NewConfig(logger logger.Logger) *viper.Viper {
	v := viper.New()
	path := "./config"
	if os.Getenv("CONFIG_PATH") != "" {
		path = os.Getenv("CONFIG_PATH")
	}
	v.AddConfigPath(path) // path to look for the config file in
	v.SetConfigName("production")
	if env := os.Getenv("ENV"); env == "development" {
		v.SetConfigName("development")
	}

	if err := v.ReadInConfig(); err != nil {
		logger.Error("Failed to load the config file", err.Error())
	}

	return v
}
