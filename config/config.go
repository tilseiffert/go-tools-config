package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Options []Option // list of used options
}

type Option struct {
	Name        string
	Default     interface{}
	HelpMessage string
}

// New() creates default configuration.
func New() Configuration {
	return Configuration{}
}

// AddOption() appends the given option to the configuration options-array
// and returns a pointer to the newly appended option.
func (c *Configuration) AddOption(o Option) *Option {
	c.Options = append(c.Options, o)
	return &c.Options[len(c.Options)-1]
}

// NewOpption() creates and adds an option with the given values to the
// options-array and returns a pointer to the new option asdf asdf adsf asdf
func (c *Configuration) NewOption(name string, defaultValue interface{}, helpMessage string) *Option {
	return c.AddOption(Option{
		Name:        name,
		Default:     defaultValue,
		HelpMessage: helpMessage,
	})
}

// NewStrongOpption() creates and adds an option with the given values and
// an empty default-value to the options-array and returns a pointer to
// the new option.
func (c *Configuration) NewStringOption(name, helpMessage string) *Option {
	return c.AddOption(Option{
		Name:        name,
		Default:     "",
		HelpMessage: helpMessage,
	})
}

// Get() retrieves the current option from viper and returns its value.
func (o *Option) Get() interface{} {
	return viper.Get(o.Name)
}

// Init() initalize viper with the given Configuration.
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
