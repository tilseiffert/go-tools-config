package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tilseiffert/go-tools-config/config"
)

func main() {

	// prepare
	configuration := config.Configuration{
		Options: []config.Option{
			{
				Name:    "name",
				Default: "Tilmann",
			},
		},
	}

	// initialize
	config.Init(configuration)

	// do your stuff
	fmt.Printf("Hello %s\n", viper.GetString(configuration.Options[0].Name))

}
