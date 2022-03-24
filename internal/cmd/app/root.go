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

func NewRootCommand() *RootCommand {
	root := &RootCommand{}
	root.Command = &cobra.Command{
		Use:           "generate [command]",
		Short:         "terraform tool",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	root.Command.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		root.log = logger.NewLogger(logger.WithName("generate-tool"))

		return nil
	}
	root.Command.AddCommand(
		generateVariableObject(root),
		generateLocalObject(root),
	)

	return root
}
