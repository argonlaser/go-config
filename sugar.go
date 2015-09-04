package config

import (
	"log"
	"os"
)

// Resolve `options` using the built-in flag and env resolvers.
func Resolve(options interface{}) error {
	c := Config{
		Options: options,
		Resolvers: []Resolver{
			&FlagResolver{Args: os.Args},
			&EnvResolver{},
		},
	}

	return c.Resolve()
}

// MustResolve `options` using the built-in flag and env resolvers.
func MustResolve(options interface{}) {
	err := Resolve(options)
	if err != nil {
		log.Fatalf("error resolving configuration: %s", err)
	}
}
