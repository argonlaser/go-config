package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEnvResolver(t *testing.T) {
	o := struct {
		Number      int           `name:"number" desc:"some number"`
		Concurrency int           `name:"concurrency" desc:"message concurrency"`
		Port        uint          `name:"port" desc:"redis port"`
		Address     string        `name:"address" desc:"redis address"`
		Addresses   []string      `name:"addresses" desc:"redis addresses"`
		Size        Bytes         `name:"size" desc:"max size"`
		Timeout     time.Duration `name:"timeout"`
	}{
		Number:      123,
		Concurrency: 10,
	}

	os.Setenv("CONCURRENCY", "5")
	os.Setenv("PORT", "3000")
	os.Setenv("ADDRESS", "0.0.0.0:3000")
	os.Setenv("ADDRESSES", "0.0.0.0:3001,0.0.0.0:3002")
	os.Setenv("SIZE", "15MiB")
	os.Setenv("TIMEOUT", "5s")

	c := Config{
		Options: &o,
		Resolvers: []Resolver{
			&EnvResolver{},
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

func TestEnvResolverPrefix(t *testing.T) {
	o := struct {
		Port    uint   `name:"port" desc:"redis port"`
		Address string `name:"address" desc:"redis address"`
	}{}

	os.Setenv("REDIS_PORT", "3000")
	os.Setenv("REDIS_ADDRESS", "0.0.0.0:3000")

	c := Config{
		Options: &o,
		Resolvers: []Resolver{
			&EnvResolver{Prefix: "redis"},
		},
	}

	err := c.Resolve()
	assert.Nil(t, err)

	assert.Equal(t, uint(3000), o.Port)
	assert.Equal(t, "0.0.0.0:3000", o.Address)
}
