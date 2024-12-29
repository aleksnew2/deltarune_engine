package console

import (
	"deltarune_engine/io/console/message"
	"fmt"
)

type Console struct {
	Settings Settings
}

// Log prints the provided arguments to the console, formatted according to the Console's settings.
// It accepts a variable number of string arguments.
func (c Console) Log(args ...string) {
	fmt.Println(format(c.Settings, message.Info, args...))
}

// Error prints a formatted error message to the console.
// The error message is formatted based on the provided error code and additional arguments.
func (c Console) Error(error_code int32, args ...string) {
	fmt.Println(formatException(c.Settings, error_code, args...))
}

// Warning logs a warning message to the console.
// It takes an error code and a variadic number of string arguments
// which are used to format the warning message.
func (c Console) Warning(error_code int32, args ...string) {
	fmt.Println(format(c.Settings, message.Warning, args...))
}
