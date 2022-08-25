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

package app

import (
	logger "github.com/AppsFlyer/go-logger"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	*cobra.Command
	log     logger.Logger
	DryRun  bool
	Verbose bool
}

func NewRootCommand(logger logger.Logger) *RootCommand {
	root := &RootCommand{}
	root.Command = &cobra.Command{
		Use:           "generate [command]",
		Short:         "terraform tool",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	root.Command.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		root.log = logger

		return nil
	}

	svc := InitTerraformGeneratorService(root.log)

	root.Command.AddCommand(
		generateTerraformFile(root, svc.GenerateModuleDefaultLocals, "terraform-locals"),
		generateTerraformFile(root, svc.GenerateModuleVariableObject, "terraform-variable"),
		generateMain(root),
		generateAllFiles(root),
	)

	return root
}
