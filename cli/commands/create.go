package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CreateCreateCommand() *cobra.Command {
	var createCommand = &cobra.Command{
		Use:     "create",
		Example: "pocketsctl create test-1",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("creating container named %s", args[0])
		},
	}
	return createCommand
}
