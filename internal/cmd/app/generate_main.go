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
	"github.com/AppsFlyer/terra-crust/internal/cmd/types"
	template_reader "github.com/AppsFlyer/terra-crust/internal/services/drivers/template_reader"
	version_control "github.com/AppsFlyer/terra-crust/internal/services/drivers/version_control"
	"github.com/spf13/cobra"
)

func generateMain(root *RootCommand) *cobra.Command {
	flags := types.TFGenerateFlags{}
	cmd := &cobra.Command{
		Use:     "terraform-main",
		Short:   "create general object terraform main file",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := root.log

			terraformSvc := InitTerraformGeneratorService(log)
			templateReader := template_reader.InitTemplateRemoteModule(log)
			gitDriver := version_control.InitGitProvider(log)

			if flags.FetchRemote && flags.MainTemplateFilePath != "" {
				log.Infof("Searching for remote modules")
				remoteModulesMap, err := templateReader.GetRemoteModulesFromTemplate(flags.MainTemplateFilePath)
				if err != nil {
					log.Error("Failed parsing remote modules from custom template", err.Error())

					return err
				}

				log.Infof("found remote modules: ", remoteModulesMap)

				if err = gitDriver.CloneModules(remoteModulesMap, flags.SourcePath); err != nil {
					log.Error("Failed cloning remote modules ", err.Error())

					return err
				}

				defer func() {
					if err = gitDriver.CleanModulesFolders(remoteModulesMap, flags.SourcePath); err != nil {
						log.Errorf("Failed to clean up some of the remote resources please make sure to clean it manually and check the error , %s", err.Error())
					}
				}()
			}

			log.Infof("remote not found ")

			if err := terraformSvc.GenerateMain(flags.SourcePath, flags.DestinationPath, flags.MainTemplateFilePath); err != nil {
				log.Error("Failed generating the terraform main file", err.Error())

				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&flags.SourcePath, "source-path", "", "Required: General module folder path that contains all the sub modules flattened")
	cmd.Flags().StringVar(&flags.DestinationPath, "destination-path", "", "Required: Destination path to write the new terraform file")
	cmd.Flags().StringVar(&flags.MainTemplateFilePath, "main-template-path", "", "Optional: Custom main template path for generated module, will take default if not provided")
	cmd.Flags().BoolVar(&flags.FetchRemote, "fetch-version-control", false, "Optional: Enable fetching of remote modules and exporting their variables in root module")
	if err := cmd.MarkFlagRequired("source-path"); err != nil {
		root.log.Error("failed to set required flag on source-path", err.Error())
	}
	if err := cmd.MarkFlagRequired("destination-path"); err != nil {
		root.log.Error("failed to set required flag on destination-path", err.Error())
	}

	return cmd
}
