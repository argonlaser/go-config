package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFlagResolver(t *testing.T) {
	o := struct {
		Number      int           `name:"number" help:"some number"`
		Concurrency int           `name:"concurrency" help:"message concurrency"`
		Port        uint          `name:"port" help:"redis port"`
		Address     string        `name:"address" help:"redis address"`
		Addresses   []string      `name:"addresses" help:"redis addresses"`
		Size        Bytes         `name:"size" help:"max size"`
		Timeout     time.Duration `name:"timeout"`
	}{
		Number:      123,
		Concurrency: 10,
	}

	c := Config{
		Options: &o,
		Resolvers: []Resolver{
			&FlagResolver{
				Args: []string{
					"program",
					"--concurrency=5",
					"--port=3000",
					"--address",
					"0.0.0.0:3000",
					"--addresses",
					"0.0.0.0:3001",
					"--addresses",
					"0.0.0.0:3002",
					"--size=15MiB",
					"--timeout=5s",
				},
			},
		},
	}

	err := c.Resolve()
	assert.Nil(t, err)

	assert.Equal(t, time.Second*5, o.Timeout)
	assert.Equal(t, 123, o.Number)
	assert.Equal(t, Bytes(15<<20), o.Size)
	assert.Equal(t, 5, o.Concurrency)
	assert.Equal(t, uint(3000), o.Port)
	assert.Equal(t, "0.0.0.0:3000", o.Address)
	assert.Equal(t, []string{"0.0.0.0:3001", "0.0.0.0:3002"}, o.Addresses)
}
