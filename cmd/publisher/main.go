package main

import (
	server "Gonverter/app"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	app := new(server.App)
	app.Run("8090")

	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
