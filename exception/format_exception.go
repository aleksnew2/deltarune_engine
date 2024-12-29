package exception

import (
	"fmt"
	"strings"
)

// format generates a formatted error message string based on the provided error code and optional arguments.
// If no arguments are provided, it returns a string in the format "deltarune_exception [error_code]".
// If arguments are provided, it combines them into a single string and returns a string in the format
// "deltarune_exception [error_code]: combined_arguments".
//
// Parameters:
//   - error_code: An int32 representing the error code.
//   - args: A variadic string slice containing additional information to include in the error message.
//
// Returns:
//   A formatted string representing the error message.
func format(error_code int32, args ...string) string {
	argsCombined := strings.Join(args, " ")
	if len(args) == 0 {
		return fmt.Sprintf("deltarune_exception [%d]", error_code)
	}
	return fmt.Sprintf("deltarune_exception [%d]: %s", error_code, argsCombined)
}
