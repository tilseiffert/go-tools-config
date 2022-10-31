package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Options []Option // list of used options
}

type Option struct {
	Name    string
	Default interface{}
}

// Init initalize viper with the given Configuration.
func Init(c Configuration) error {

	// set default for each option
	for _, v := range c.Options {
		// if v.Default == nil {
		// 	continue
		// }
		viper.SetDefault(v.Name, v.Default)
	}

	return nil
}
