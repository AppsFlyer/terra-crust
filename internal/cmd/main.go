package main

import (
	"context"
	"os"

	logger "github.com/AppsFlyer/go-logger"
	"github.com/AppsFlyer/terra-crust/internal/cmd/app"
)

func main() {
	log := logger.NewSimple()
	if err := app.NewRootCommand(log).Execute(); err != nil {
		if err == context.Canceled {
			os.Exit(0)
		}

		log.Error("Error executing command. Exiting.", err.Error())
		os.Exit(1)
	}
}
