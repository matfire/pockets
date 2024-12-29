package commands

import (
	"github.com/charmbracelet/huh"
	"github.com/matfire/pockets/cli/config"
	"github.com/matfire/pockets/cli/handlers/docker"
	"github.com/spf13/cobra"
)

func CreateCreateCommand(config *config.App) *cobra.Command {
	var createCommand = &cobra.Command{
		Use:     "create",
		Example: "pocketsctl create test-1",
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var name string
			if len(args) > 0 {
				name = args[0]
			} else {
				huh.NewInput().Title("what's the name of the container?").Value(&name).Run()
			}
			docker.Create(config, name)
		},
	}
	return createCommand
}
