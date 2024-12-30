package console

import (
	"fmt"
	"os"

	"github.com/aleksnew2/deltarune_engine/exception"
	"github.com/aleksnew2/deltarune_engine/io/console/message"
)

// Console represents a console with specific settings.
type Console struct {
	Settings Settings
}

// InitConsole initializes and returns a new Console instance with default settings.
// The default settings are:
//
// 	- IsWithTimestamp: false
//  - IsWithExceptionCode: true
//  - ErrorExceptions: an empty slice of exceptions
//  - TerminateOnError: false
func Init() Console {
	return Console{
		Settings: Settings{
			IsWithTimestamp:     false,
			IsWithExceptionCode: true,
			ErrorExceptions:     make([]exception.Exception, 0),
			TerminateOnError:    false,
		},
	}
}

// Log prints the provided arguments to the console with the specified settings and message type.
// It accepts a variable number of string arguments.
//
// Parameters:
//   - args: A variable number of string arguments to be logged.
func (c Console) Log(args ...string) {
	fmt.Println(format(c.Settings, message.Info, args...))
}

// Error prints a formatted error message to the console based on the provided error code and optional arguments.
// If the Console's TerminateOnError setting is true, the program will exit with status code 0 after printing.
//
// Parameters:
//   - error_code: An integer error code that identifies the type of error
//   - args: Optional variadic string arguments used to format the error message
func (c Console) Error(error_code int32, args ...string) {
	fmt.Println(formatException(c.Settings, error_code, args...))

	if c.Settings.TerminateOnError {
		os.Exit(0)
	}
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
