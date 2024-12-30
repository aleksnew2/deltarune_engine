package console

import (
	"fmt"
	"github.com/aleksnew2/deltarune_engine/io/console/message"
)

type Console struct {
	Settings Settings
}

// Log prints the provided arguments to the console with the specified settings and message type.
// It accepts a variable number of string arguments.
//
// Parameters:
//   - args: A variable number of string arguments to be logged.
func (c Console) Log(args ...string) {
	fmt.Println(format(c.Settings, message.Info, args...))
}

// Error logs an error message to the console based on the provided error code and arguments.
// The error message is formatted using the Console's settings.
//
// Parameters:
//   - error_code: An int32 representing the error code.
//   - args: A variadic list of strings that provide additional context or details for the error message.
func (c Console) Error(error_code int32, args ...string) {
	fmt.Println(formatException(c.Settings, error_code, args...))
}

// Warning logs a warning message to the console.
// It takes an error code and a variadic number of string arguments
// which are formatted and printed according to the console settings.
//
// Parameters:
//   - error_code: An int32 representing the error code.
//   - args: A variadic number of string arguments to be included in the warning message.
func (c Console) Warning(error_code int32, args ...string) {
	fmt.Println(format(c.Settings, message.Warning, args...))
}
