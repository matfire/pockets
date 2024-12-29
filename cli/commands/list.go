package commands

import (
	"github.com/matfire/pockets/cli/config"
	"github.com/matfire/pockets/cli/handlers/docker"
	"github.com/spf13/cobra"
)

func CreateListCommand(config *config.App) *cobra.Command {
	var listCommand = &cobra.Command{
		Use:     "list",
		Example: "pocketsctl list",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			docker.List(config)
		},
	}
	return listCommand
}
