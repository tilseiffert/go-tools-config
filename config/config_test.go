package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {

	// setup

	config := Configuration{
		Options: []Option{
			{
				Name:    "ONE",
				Default: 1,
			},
			{
				Name:    "NIL",
				Default: nil,
			},
		},
	}

	// prepare

	viper.Reset()

	// run command

	if assert.Nil(t, Init(config)) {

		// run tests

		t.Run("Check options and their default values", func(t *testing.T) {

			for _, v := range config.Options {
				assert.Equal(t, v.Default, viper.Get(v.Name))
			}

		})

	}

}
