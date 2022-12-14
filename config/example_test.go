package config_test

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tilseiffert/go-tools-config/config"
)

func ExampleInit() {

	// preparation for test environment
	viper.Reset() // not needed for productive usage

	// prepare
	conf := config.New()

	optionName := conf.NewOption("NAME", "Tilmann", false, "")

	// initialize
	config.Init(conf)
	viper.SetTypeByDefaultValue(true)

	// do your stuff
	fmt.Printf("Hello %s\n", optionName.Get())

	// Output: Hello Tilmann
}
