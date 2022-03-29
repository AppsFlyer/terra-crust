package main

import (
	"context"
	"os"

	logger "gitlab.appsflyer.com/go/af-go-logger/v1"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/cmd/app"
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
