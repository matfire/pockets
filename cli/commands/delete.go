package commands

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/matfire/pockets/cli/config"
	"github.com/matfire/pockets/cli/handlers/docker"
	"github.com/spf13/cobra"
)

func CreateDeleteCommand(config *config.App) *cobra.Command {
	var createCommand = &cobra.Command{
		Use:     "delete",
		Example: "pocketsctl start test-1",
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var name string
			var confirmed bool
			if len(args) > 0 {
				name = args[0]
			} else {
				huh.NewInput().Title("what's the id of the container?").Value(&name).Run()
			}
			huh.NewConfirm().Title("Are you sure you want to delete this container?").Value(&confirmed).Run()
			if confirmed {
				docker.Delete(config, name)
			} else {
				fmt.Print("deletion request ignored")
			}
		},
	}
	return createCommand
}
