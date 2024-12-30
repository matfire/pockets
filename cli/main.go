package main

import (
	"github.com/matfire/pockets/cli/commands"
	"github.com/matfire/pockets/cli/config"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("ENDPOINT", "http://127.0.0.1:3000")
	viper.SetConfigName("pockets")
	viper.AddConfigPath("$HOME/.config/pockets")
	viper.AddConfigPath(".")
	// you can pass in the endpoint value using PKTS_ENDPOINT="<value>"
	viper.SetEnvPrefix("PKTS")
	_ = viper.ReadInConfig()

	var config config.App
	config.Endpoint = viper.GetString("ENDPOINT")

	commands.RootCmd.AddCommand(commands.CreateCreateCommand(&config), commands.CreateListCommand(&config), commands.CreateStopCommand(&config), commands.CreateStartCommand(&config), commands.CreateDeleteCommand(&config))
	commands.RootCmd.Execute()
}
