package config_test

import (
	"github.com/tj/go-config"
)

type Options struct {
	Concurrency uint   `desc:"max in-flight messages"`
	LogLevel    string `desc:"log level"`
}

// ExampleResolve illustrates the simplest way to use go-config. Using
// the MustResolve function pre-configures the flag and env resolvers for
// the average use-case.
func ExampleResolve() {
	options := &Options{
		Concurrency: 5,
		LogLevel:    "info",
	}

	err := config.Resolve(options)
	if err != nil {
		panic(err)
	}
}
