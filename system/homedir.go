package system

import (
	"deltarune_engine/exception"
	"os"
)

// Homedir returns the current user's home directory path.
//
// The function uses os.UserHomeDir() to retrieve the home directory path.
// If the operation fails, it returns an empty string and a wrapped error.
func Homedir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", exception.CreateException(-1, err.Error())
	}

	return homeDir, nil
}
