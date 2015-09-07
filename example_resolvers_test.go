package config_test

import (
	"log"
	"os"
	"time"

	"github.com/tj/go-config"
)

type ResolverOptions struct {
	Timeout     time.Duration `help:"message timeout"`
	Concurrency uint          `help:"max in-flight messages"`
	CacheSize   config.Bytes  `help:"cache size in bytes"`
	BatchSize   uint          `help:"batch size" validate:"min=1,max=1000"`
	LogLevel    string        `help:"set the log severity" from:"env,flag"`
}

// ExampleResolvers illustrates how you may initialize a Config
// struct in order to provide custom resolvers for more flexibility.
func Example_resolvers() {
	options := &ResolverOptions{
		Timeout:     5 * time.Second,
		Concurrency: 5,
		CacheSize:   config.ParseBytes("150mb"),
		BatchSize:   1000,
		LogLevel:    "info",
	}

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
