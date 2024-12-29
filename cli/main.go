package main

import (
	"github.com/matfire/pockets/cli/commands"
)

func main() {
	commands.RootCmd.AddCommand(commands.CreateCreateCommand(), commands.CreateListCommand())
	commands.RootCmd.Execute()
}
