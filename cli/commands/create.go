package commands

import (
	"github.com/matfire/pockets/cli/config"
	"github.com/matfire/pockets/cli/handlers/docker"
	"github.com/spf13/cobra"
)

func CreateCreateCommand(config *config.App) *cobra.Command {
	var createCommand = &cobra.Command{
		Use:     "create",
		Example: "pocketsctl create test-1",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			docker.Create(config, args[0])
		},
	}
	return createCommand
}
