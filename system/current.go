package system

import (
	"github.com/aleksnew2/deltarune_engine/exception"
	"os"
)

// Current returns the hostname of the current operating system.
// It retrieves the hostname using the os.Hostname() function.
// If there is an error retrieving the hostname, it returns an error.
// Otherwise, it returns the hostname prefixed with "OS: ".
func Current() (string, error) {
	name, err := os.Hostname()
	if err != nil {
		return "", exception.CreateException(-1, err.Error())
	}

	return "OS: " + name, nil
}
