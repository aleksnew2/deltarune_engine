package system

import (
	"deltarune_engine/exception"
	"os/user"
)

// GetCurrentUserName retrieves the username of the currently logged-in user.
// It uses the os/user package to get the current user information.
func GetCurrentUserName() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", exception.CreateException(-1, err.Error())
	}
	return currentUser.Username, nil
}
