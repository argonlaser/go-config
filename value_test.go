package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolValueSet(t *testing.T) {
	{
		v := boolValue(false)
		err := v.Set("1")
		assert.Nil(t, err)
		assert.Equal(t, "true", v.String())
	}

	{
		v := boolValue(false)
		err := v.Set("t")
		assert.Nil(t, err)
		assert.Equal(t, "true", v.String())
	}

	{
		v := boolValue(false)
		err := v.Set("true")
		assert.Nil(t, err)
		assert.Equal(t, "true", v.String())
	}

	{
		v := boolValue(true)
		err := v.Set("0")
		assert.Nil(t, err)
		assert.Equal(t, "false", v.String())
	}

	{
		v := boolValue(true)
		err := v.Set("f")
		assert.Nil(t, err)
		assert.Equal(t, "false", v.String())
	}

	{
		v := boolValue(true)
		err := v.Set("false")
		assert.Nil(t, err)
		assert.Equal(t, "false", v.String())
	}
}

func TestIntValueSet(t *testing.T) {
	{
		v := intValue(123)
		err := v.Set("50")
		assert.Nil(t, err)
		assert.Equal(t, "50", v.String())
	}

	{
		v := intValue(123)
		err := v.Set("asdf")
		assert.NotNil(t, err)
	}
}

func TestUintValueSet(t *testing.T) {
	{
		v := uintValue(123)
		err := v.Set("50")
		assert.Nil(t, err)
		assert.Equal(t, "50", v.String())
	}

	{
		v := uintValue(123)
		err := v.Set("-15")
		assert.NotNil(t, err)
	}
}

func TestFloatValueSet(t *testing.T) {
	{
		v := floatValue(123)
		err := v.Set("50.99")
		assert.Nil(t, err)
		assert.Equal(t, "50.99", v.String())
	}

	{
		v := floatValue(123)
		err := v.Set("asdf")
		assert.NotNil(t, err)
	}
}

func TestStringValueSet(t *testing.T) {
	{
		v := stringValue("something")
		err := v.Set("hello")
		assert.Nil(t, err)
		assert.Equal(t, "hello", v.String())
	}
}

func TestStringsValueSet(t *testing.T) {
	{
		v := stringsValue{}
		assert.Nil(t, v.Set("hello"))
		assert.Nil(t, v.Set("world"))
		assert.Equal(t, "hello,world", v.String())
	}

	{
		v := stringsValue{}
		assert.Nil(t, v.Set("foo,bar"))
		assert.Nil(t, v.Set("baz"))
		assert.Equal(t, "foo,bar,baz", v.String())
	}
}

func TestBytes_Set(t *testing.T) {
	{
		v := Bytes(123)
		err := v.Set("50mb")
		assert.Nil(t, err)
		assert.Equal(t, "50MB", v.String())
	}

	{
		v := Bytes(123)
		err := v.Set("asdf")
		assert.NotNil(t, err)
	}
}
