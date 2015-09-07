package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestField_Name(t *testing.T) {
	a := struct {
		Hello string
	}{"hello"}

	fa := field{
		value: reflect.ValueOf(a).Field(0),
		field: reflect.ValueOf(a).Type().Field(0),
	}

	b := struct {
		World string
	}{"world"}

	fb := field{
		value:  reflect.ValueOf(b).Field(0),
		field:  reflect.ValueOf(b).Type().Field(0),
		parent: &fa,
	}

	assert.Equal(t, "hello", fa.Name())
	assert.Equal(t, "hello-world", fb.Name())
}
