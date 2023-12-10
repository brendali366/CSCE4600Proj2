package builtins

import (
	"os"
)

// bye exits the program
func ByeCommand(args ...string) {
	os.Exit(0)
}
