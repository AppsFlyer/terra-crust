package main

import (
	"context"
	"os"

	logger "github.com/AppsFlyer/go-logger"
	"github.com/AppsFlyer/terra-crust/internal/cmd/app"
)

func main() {
	log := logger.NewLogger(logger.WithName("terraform-generate-tool"))
	if err := app.NewRootCommand(log).Execute(); err != nil {
		if err == context.Canceled {
			os.Exit(0)
		}

		log.ErrorWithError("Error executing command. Exiting.", err)
		os.Exit(1)
	}
}
