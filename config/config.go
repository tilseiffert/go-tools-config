package config

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Options               []Option // list of used options
	SetTypeByDefaultValue bool
	SetAutomaticEnv       bool
}

type Option struct {
	Name          string
	Default       interface{}
	HelpMessage   string
	CheckNotEmpty bool // only implemented for string and int
}

type ErrConfigOptionsEmpty struct {
	FailedOptions []string
}

// define Error() method on the struct
func (e ErrConfigOptionsEmpty) Error() string {

	var msg bytes.Buffer
	msg.WriteString(fmt.Sprintf("there were %d unexpected empty options: ", len(e.FailedOptions)))

	for k, v := range e.FailedOptions {
		if k != 0 {
			msg.WriteString(", ")
		}

		msg.WriteString(v)
	}

	return msg.String()
}

// New() creates default configuration:
// - SetTypeByDefaultValue: true
// - SetAutomaticEnv: true,
func New() Configuration {
	return Configuration{
		Options:               []Option{},
		SetTypeByDefaultValue: true,
		SetAutomaticEnv:       true,
	}
}

// AddOption() appends the given option to the configuration options-array
// and returns a pointer to the newly appended option.
func (c *Configuration) AddOption(o Option) *Option {
	c.Options = append(c.Options, o)
	return &c.Options[len(c.Options)-1]
}

// NewOpption() creates and adds an option with the given values to the
// options-array and returns a pointer to the new option asdf asdf adsf asdf
func (c *Configuration) NewOption(name string, defaultValue interface{}, checkNotEmpty bool, helpMessage string) *Option {
	return c.AddOption(Option{
		Name:          name,
		Default:       defaultValue,
		CheckNotEmpty: checkNotEmpty,
		HelpMessage:   helpMessage,
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

// GetString() retrieves the current option from viper and returns its value
// as string.
func (o *Option) GetString() string {
	return viper.GetString(o.Name)
}

// GetInt() retrieves the current option from viper and returns its value
// as int.
func (o *Option) GetInt() int {
	return viper.GetInt(o.Name)
}

// GetBool() retrieves the current option from viper and returns its value
// as bool.
func (o *Option) GetBool() bool {
	return viper.GetBool(o.Name)
}

// GetFloat64() retrieves the current option from viper and returns its value
// as float64.
func (o *Option) GetFloat64() float64 {
	return viper.GetFloat64(o.Name)
}

// Init() initalize viper with the given Configuration.
func Init(c Configuration) error {

	optionsEmpty := ErrConfigOptionsEmpty{}

	// set default for each option
	for _, v := range c.Options {
		viper.SetDefault(v.Name, v.Default)
	}

	if c.SetAutomaticEnv {
		viper.AutomaticEnv()
	}

	if c.SetTypeByDefaultValue {
		viper.SetTypeByDefaultValue(true)
	}

	for _, v := range c.Options {
		if !v.CheckNotEmpty {
			continue
		}

		switch v.Default.(type) {
		case int:
			if v.Get() == 0 {
				optionsEmpty.FailedOptions = append(optionsEmpty.FailedOptions, v.Name)
			}
		case string:
			if v.Get() == "" {
				optionsEmpty.FailedOptions = append(optionsEmpty.FailedOptions, v.Name)
			}
		default:
			if v.Get() == nil {
				optionsEmpty.FailedOptions = append(optionsEmpty.FailedOptions, v.Name)
			}
		}
	}

	if len(optionsEmpty.FailedOptions) > 0 {
		return optionsEmpty
	}

	return nil
}
