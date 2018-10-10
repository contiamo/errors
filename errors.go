package errors

import (
	"github.com/pkg/errors"
)

/**
 * The following are just wrappers around github.com/pkg/errors to prevent import headache
 */

// Wrap annotates the error with a message
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

// Wrapf annotates the errors with a formatted message
func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

// WithStack annotates the error with a stack trace
func WithStack(err error) error {
	return errors.WithStack(err)
}
