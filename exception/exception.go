// Package exception provides utilities for creating and handling custom exceptions.
//
// The Exception type is a type alias for the built-in error type, allowing for
// more descriptive error handling in the program.
//
// The CreateException function allows for the creation of a new Exception with
// a specific error code and optional additional context or details.
package exception

import "errors"

// Exception is a type alias for the built-in error type.
// It can be used to represent an error condition in the program.
type Exception error

// CreateException creates a new Exception with the given error code and optional arguments.
// The error code is an int32 representing the specific error.
// The args parameter is a variadic string slice that can be used to provide additional context or details for the error.
//
// Parameters:
//   - error_code: An int32 representing the specific error code.
//   - args: A variadic string slice for additional context or details.
//
// Returns:
//   - Exception: A new Exception instance with the formatted error message.
func CreateException(error_code int32, args ...string) Exception {
	return errors.New(format(error_code, args...))
}
