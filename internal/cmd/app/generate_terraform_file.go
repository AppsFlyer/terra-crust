package app

import (
	"fmt"

	"github.com/AppsFlyer/terra-crust/internal/cmd/types"
	"github.com/spf13/cobra"
)

func generateTerraformFile(root *RootCommand, f func(modulesFilePath string, destinationPath string) error, short string) *cobra.Command {
	var flags types.TFGenerateFlags = types.TFGenerateFlags{}
	cmd := &cobra.Command{
		Use:     short,
		Short:   fmt.Sprintf("create general object %s file", short),
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := root.log

			if err := f(flags.SourcePath, flags.DestinationPath); err != nil {
				log.Error("Failed generating the terraform locals file", err.Error())

				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&flags.DestinationPath, "destination-path", "", "Required: Destination path to write the new terraform file")
	cmd.Flags().StringVar(&flags.SourcePath, "source-path", "", "Required:  General module folder path that contains all the sub modules flattened")
	if err := cmd.MarkFlagRequired("source-path"); err != nil {
		root.log.Error("failed to set required flag on source-path", err.Error())
	}
	if err := cmd.MarkFlagRequired("destination-path"); err != nil {
		root.log.Error("failed to set required flag on destination-path", err.Error())
	}

	return cmd
}
