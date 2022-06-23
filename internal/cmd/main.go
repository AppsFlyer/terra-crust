// Copyright 2022 AppsFlyer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
