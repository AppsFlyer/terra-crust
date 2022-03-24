package app

import (
	"os"

	"github.com/spf13/viper"
	logger "gitlab.appsflyer.com/go/af-go-logger/v1"
)

func NewConfig(logger logger.Logger) *viper.Viper {
	v := viper.New()
	v.AddConfigPath("./config") // path to look for the config file in
	v.SetConfigName("production")
	if env := os.Getenv("ENV"); env == "development" {
		v.SetConfigName("development")
	}

	if err := v.ReadInConfig(); err != nil {
		logger.ErrorWithError("Failed to load the config file", err)
	}

	return v
}
