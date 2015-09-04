package config

import (
	"os"
	"strings"
)

// EnvResolver resolves configuration from environment variables.
//
// For example LogLevel field would become LOG_LEVEL=error.
type EnvResolver struct {
	Prefix string // Prefix optionally applied to each lookup
}

// Name implementation.
func (e *EnvResolver) Name() string {
	return "env"
}

// Setup implementation (temporary noop).
func (e *EnvResolver) Setup() error {
	return nil
}

// Field implementation normalizing the field name
// and performing coercion to the field type.
func (e *EnvResolver) Field(field Field) error {
	name := field.Name()
	s := os.Getenv(e.envize(name))

	if s == "" {
		return ErrFieldNotFound
	}

	return field.Value().Set(s)
}

// Resolve implementation (temporary noop).
func (f *EnvResolver) Resolve() error {
	return nil
}

// Normalize `name` with prefix support.
func (f *EnvResolver) envize(name string) string {
	if f.Prefix != "" {
		return f.normalize(f.Prefix) + "_" + f.normalize(name)
	}

	return f.normalize(name)
}

// Normalize `name`.
func (f *EnvResolver) normalize(name string) string {
	return strings.ToUpper(strings.Replace(name, "-", "_", -1))
}
