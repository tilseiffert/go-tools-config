package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func testInitAndOptions(t *testing.T, conf Configuration, options []Option) {
	t.Run("test init and options", func(t *testing.T) {

		// run Init
		if !assert.Nil(t, Init(conf)) {
			return
		}

		// run tests
		for _, v := range conf.Options {
			assert.Equal(t, v.Default, viper.Get(v.Name), "retrieve value of option through viper directly")
		}

		for _, v := range options {
			assert.Equal(t, v.Default, v.Get(), "retrieve value of option through getta")
		}
	})
}

func TestInit(t *testing.T) {

	// prepare test
	viper.Reset()

	// setup
	options := []Option{
		{
			Name:    "ONE",
			Default: 1,
		},
		{
			Name:    "NIL",
			Default: nil,
		},
		{
			Name:    "STRING",
			Default: "string_0987sovho32!§$%",
		},
		{
			Name:    "emptyString",
			Default: "",
		},
	}

	conf := Configuration{
		Options: options,
	}

	// run test
	testInitAndOptions(t, conf, options)
}

func TestConfigurationAddOption(t *testing.T) {

	// prepare test
	viper.Reset()

	// setup
	conf := Configuration{}

	options := []Option{
		{
			Name:    "A",
			Default: "a",
		},
		{
			Name:    "B",
			Default: "b",
		},
	}

	optionA := conf.AddOption(options[0])
	optionB := conf.AddOption(options[1])

	// run test
	testInitAndOptions(t, conf, options)

	t.Run("compare options", func(t *testing.T) {
		compare := func(t *testing.T, o1 Option, o2 Option) {
			assert.NotPanics(t, func() {
				assert.True(t, cmp.Equal(o1, o2), "compare two option-structs")
			})
		}

		compare(t, *optionA, options[0])
		compare(t, *optionB, options[1])
	})
}

func TestConfigurationNewOption(t *testing.T) {

	// prepare test
	viper.Reset()

	// setup
	var options []Option
	conf := Configuration{}

	optionA := conf.NewOption("A", "a", "helpMessage")
	optionB := conf.NewOption("B", 2, "helpMessage")
	optionC := conf.NewOption("C", nil, "helpMessage")

	options = append(options, *optionA)
	options = append(options, *optionB)
	options = append(options, *optionC)

	// run test

	assert.Equal(t, "A", optionA.Name)
	assert.Equal(t, "a", optionA.Default)

	assert.Equal(t, "B", optionB.Name)
	assert.Equal(t, 2, optionB.Default)

	assert.Equal(t, "C", optionC.Name)
	assert.Equal(t, nil, optionC.Default)

	testInitAndOptions(t, conf, options)
}

func TestConfigurationNewStringOption(t *testing.T) {

	// prepare test
	viper.Reset()

	// setup
	var options []Option
	conf := Configuration{}

	optionA := conf.NewStringOption("A", "helpMessage")
	optionB := conf.NewStringOption("B", "helpMessage")

	options = append(options, *optionA)
	options = append(options, *optionB)

	// run test

	assert.Equal(t, "A", optionA.Name)
	assert.Equal(t, "", optionA.Default)

	assert.Equal(t, "B", optionB.Name)
	assert.Equal(t, "", optionB.Default)

	testInitAndOptions(t, conf, options)
}
