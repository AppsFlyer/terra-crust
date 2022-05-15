package app

import (
	"github.com/spf13/cobra"
	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/cmd/types"
)

func generateMain(root *RootCommand) *cobra.Command {
	var flags types.TFGenerateFlags = types.TFGenerateFlags{}
	cmd := &cobra.Command{
		Use:     "terraform-main",
		Short:   "create general object terraform main file",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := root.log.WithName("generate-main-file")

			terraformSvc := InitTerraformGeneratorService(log)

			if err := terraformSvc.GenerateMain(flags.SourcePath, flags.DestinationPath, flags.MainTemplateFilePath); err != nil {
				log.ErrorWithError("Failed generating the terraform main file", err)

				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&flags.SourcePath, "source-path", "", "Required: General module folder path that contains all the sub modules flattened")
	cmd.Flags().StringVar(&flags.DestinationPath, "destination-path", "", "Required: Destination path to write the new terraform file")
	cmd.Flags().StringVar(&flags.MainTemplateFilePath, "main-template-path", "", "Optional: Custom main template path for generated module, will take default if not provided")
	if err := cmd.MarkFlagRequired("source-path"); err != nil {
		root.log.ErrorWithError("failed to set required flag on source-path", err)
	}
	if err := cmd.MarkFlagRequired("destination-path"); err != nil {
		root.log.ErrorWithError("failed to set required flag on destination-path", err)
	}

	return cmd
}
