package system

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/aleksnew2/deltarune_engine/io/console"
)

var cns console.Console = console.Init()

func isAdmin() bool {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", "net", "session")
		err := cmd.Run()
		return err == nil
	}
	return os.Getuid() == 0
}

// RequestAdminRights attempts to re-run the current executable with administrative privileges.
// On Windows, it uses the "runas" command to run the executable as the Administrator user.
// On Unix-like systems, it uses "sudo" to run the executable with elevated privileges.
// If the command fails to execute, it logs an error using the cns.Error function.
func RequestAdminRights() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("runas", "/user:Administrator", os.Args[0])
		err := cmd.Run()
		if err != nil {
			cns.Error(-1, err.Error())
		}
	} else {
		cmd := exec.Command("sudo", os.Args[0])
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			cns.Error(-1, err.Error())
		}
	}
}
