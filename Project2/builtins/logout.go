package builtins

import (
	"os"
)

// logout exits the terminal
func LogoutCommand(args ...string) {
	os.Exit(1)
}
