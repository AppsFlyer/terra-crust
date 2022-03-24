package app

import (
	"github.com/spf13/cobra"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/cmd/types"
)

func generateLocalObject(root *RootCommand) *cobra.Command {
	var flags types.TFGenerateFlags = types.TFGenerateFlags{}
	log := root.log.WithName("kafka-create")
	cmd := &cobra.Command{
		Use:     "terraform-locals",
		Short:   "create general object terraform locals file",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := root.log.WithName("generate-locals-object")

			terraformSvc := InitTerraformGeneratorService(log)

			if err := terraformSvc.GenerateModuleDefaultLocals(flags.SourcePath, flags.DestinationPath); err != nil {
				log.ErrorWithError("Failed generating the terraform locals file", err)

				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&flags.SourcePath, "source-path", "", "General module folder path that contains all the sub modules flattened")
	cmd.Flags().StringVar(&flags.DestinationPath, "destination-path", "", "Destination path to write the new terraform file")
	if err := cmd.MarkFlagRequired("source-path"); err != nil {
		log.ErrorWithError("failed to set required flag on source-path", err)
	}
	if err := cmd.MarkFlagRequired("destination-path"); err != nil {
		log.ErrorWithError("failed to set required flag on destination-path", err)
	}

	return cmd
}
