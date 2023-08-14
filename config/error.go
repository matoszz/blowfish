// Package config handles config and config errors
package config

import (
	"fmt"
)

// ArgumentError is returned when invalid arguments are provided
type ArgumentError struct {
	Argument string
	Value    string
}

// Error returns the ArgumentError in string format
func (e *ArgumentError) Error() string {
	return fmt.Sprintf("invalid %s provided: %s", e.Argument, e.Value)
}

// NewArgumentError returns dumb errors
func NewArgumentError(arg string, value string) *ArgumentError {
	return &ArgumentError{
		Argument: arg,
		Value:    value,
	}
}
