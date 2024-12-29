package console

import "deltarune_engine/exception"

// Settings holds configuration options for the console I/O operations.
// It includes options for timestamping, exception handling, and error termination.
//
// Fields:
// - IsWithTimestamp: If true, timestamps will be included in the console output.
// - IsWithExceptionCode: If true, exception codes will be included in the console output.
// - TerminateOnError: If true, the application will terminate on encountering an error.
// - ErrorExceptions: A list of exceptions that are considered non-critical and will not cause termination.
type Settings struct {
	IsWithTimestamp     bool
	IsWithExceptionCode bool
	TerminateOnError    bool
	ErrorExceptions     []exception.Exception
}
