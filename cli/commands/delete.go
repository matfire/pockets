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
			confirmed, err := cmd.Flags().GetBool("force")
			if err != nil {
				fmt.Printf("%v", err)
				panic("error while getting flag")
			}
			if len(args) > 0 {
				name = args[0]
			} else {
				huh.NewInput().Title("what's the id of the container?").Value(&name).Run()
			}
			if !confirmed {
				huh.NewConfirm().Title("Are you sure you want to delete this container?").Value(&confirmed).Run()
			}
			if confirmed {
				docker.Delete(config, name)
			} else {
				fmt.Print("deletion request ignored")
			}
		},
	}
	createCommand.Flags().Bool("force", false, "")
	return createCommand
}
