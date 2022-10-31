package main

import (
	"fmt"

	"github.com/tilseiffert/go-tools-config/config"
)

func main() {

	fmt.Println("Hello World")

	config.Init(config.Configuration{
		Options: []config.Option{
			{
				Name:    "ONE",
				Default: 1,
			},
		},
	})

}
