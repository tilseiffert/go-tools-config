package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tilseiffert/go-tools-config/config"
)

func main() {

	// prepare
	conf := config.Configuration{
		Options: []config.Option{
			{
				Name:    "name",
				Default: "Tilmann",
			},
		},
	}

	// initialize
	config.Init(conf)

	// do your stuff
	fmt.Printf("Hello %s\n", viper.GetString(conf.Options[0].Name))

}
