package app

import (
	"os"

	"github.com/spf13/viper"
	logger "gitlab.appsflyer.com/go/af-go-logger/v1"
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
		logger.ErrorWithError("Failed to load the config file", err)
	}

	return v
}
