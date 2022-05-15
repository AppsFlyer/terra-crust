package app

import (
	"github.com/spf13/cobra"
	logger "gitlab.appsflyer.com/go/af-go-logger/v1"
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
		root.log = logger.WithName("generate-tool")

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
