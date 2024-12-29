package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "pocketsctl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello world!")
	},
}
