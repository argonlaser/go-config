package config_test

import (
	"log"
	"os"

	"github.com/tj/go-config"
)

type NestedOptions struct {
	LogLevel string `help:"set the log severity"`
	NSQ      struct {
		Address     string   `help:"nsqd address"`
		Lookup      []string `help:"nsqlookupd addresses"`
		MaxInFlight int      `help:"nsqd max in flight messages"`
	}
}

// ExampleNested illustrates how nested structs may be used. In this
// example --nsq-address and --nsq-max-in-flight flags would be
// available, as well as NSQ_ADDRESS and NSQ_MAX_IN_FLIGHT.
func Example_nested() {
	options := &NestedOptions{}

	c := config.Config{
		Options: options,
		Resolvers: []config.Resolver{
			&config.FlagResolver{Args: os.Args},
			&config.EnvResolver{},
		},
	}

	err := c.Resolve()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
