package utils

import "github.com/spf13/viper"

func GetConfig() *App {
	adminUser := viper.GetString("ADMIN_USER")
	adminPassword := viper.GetString("ADMIN_PASSWORD")
	app := App{
		AdminUser:     adminUser,
		AdminPassword: adminPassword,
	}
	return &app
}
