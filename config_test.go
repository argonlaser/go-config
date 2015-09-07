package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigResolveValidation(t *testing.T) {
	o := struct {
		Concurrency int `name:"concurrency" help:"some number" validate:"min=1,max=5"`
	}{}

	c := Config{
		Options: &o,
		Resolvers: []Resolver{
			&FlagResolver{
				Args: []string{
					"program",
					"--concurrency=0",
				},
			},
		},
	}

	err := c.Resolve()
	assert.EqualError(t, err, `Concurrency: less than min`)
}

func TestConfigResolveDefaultName(t *testing.T) {
	o := struct {
		MaxInFlight int
	}{}

	c := Config{
		Options: &o,
		Resolvers: []Resolver{
			&FlagResolver{
				Args: []string{
					"program",
					"--max-in-flight=5",
				},
			},
		},
	}

	err := c.Resolve()
	assert.NoError(t, err)
	assert.Equal(t, 5, o.MaxInFlight)
}

func TestConfigResolveTagName(t *testing.T) {
	o := struct {
		MaxInFlight int `name:"concurrency"`
	}{}

	c := Config{
		Options: &o,
		Resolvers: []Resolver{
			&FlagResolver{
				Args: []string{
					"program",
					"--concurrency=5",
				},
			},
		},
	}

	err := c.Resolve()
	assert.NoError(t, err)
	assert.Equal(t, 5, o.MaxInFlight)
}

type Redis struct {
	Host string
	Port int
}

func TestConfigResolveNested(t *testing.T) {
	o := struct {
		NSQ struct {
			Messages struct {
				MaxInFlight int
			}
		}
		Redis Redis
	}{
		Redis: Redis{
			Host: "localhost",
			Port: 6379,
		},
	}

	c := Config{
		Options: &o,
		Resolvers: []Resolver{
			&FlagResolver{
				Args: []string{
					"program",
					"--nsq-messages-max-in-flight=5",
					"--redis-host=0.0.0.0",
					"--redis-port=3000",
				},
			},
		},
	}

	err := c.Resolve()
	assert.NoError(t, err)
	assert.Equal(t, 5, o.NSQ.Messages.MaxInFlight)
	assert.Equal(t, "0.0.0.0", o.Redis.Host)
	assert.Equal(t, 3000, o.Redis.Port)
}

type Foo struct {
	Bar *Bar
}

type Bar struct {
	Baz string
}

func TestConfigResolveNestedPointers(t *testing.T) {
	o := struct {
		Foo *Foo
	}{
		Foo: &Foo{
			Bar: &Bar{
				Baz: "hello",
			},
		},
	}

	c := Config{
		Options: &o,
		Resolvers: []Resolver{
			&FlagResolver{
				Args: []string{
					"program",
					"--foo-bar-baz=world",
				},
			},
		},
	}

	err := c.Resolve()
	assert.NoError(t, err)
	assert.Equal(t, "world", o.Foo.Bar.Baz)
}
