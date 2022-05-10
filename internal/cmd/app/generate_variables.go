package app

import (
	"github.com/spf13/cobra"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/cmd/types"
)

func generateVariableObject(root *RootCommand) *cobra.Command {
	var flags types.TFGenerateFlags = types.TFGenerateFlags{}
	cmd := &cobra.Command{
		Use:     "terraform-variable",
		Short:   "create general object terraform variable file",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := root.log.WithName("generate-variable-object")

			terraformSvc := InitTerraformGeneratorService(log)

			if err := terraformSvc.GenerateModuleVariableObject(flags.SourcePath, flags.DestinationPath); err != nil {
				log.ErrorWithError("Failed generating the terraform variable file", err)

				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&flags.SourcePath, "source-path", "", "Required: General module folder path that contains all the sub modules flattened")
	cmd.Flags().StringVar(&flags.DestinationPath, "destination-path", "", "Required: Destination path to write the new terraform file")
	if err := cmd.MarkFlagRequired("source-path"); err != nil {
		root.log.ErrorWithError("failed to set required flag on source-path", err)
	}
	if err := cmd.MarkFlagRequired("destination-path"); err != nil {
		root.log.ErrorWithError("failed to set required flag on destination-path", err)
	}

	return cmd
}
