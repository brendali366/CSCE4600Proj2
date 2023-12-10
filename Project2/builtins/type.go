package builtins

import (
	"fmt"
	"os/exec"
)

func TypeCommand(args ...string) {
	cmd := exec.Command("type", args[0])

	output, err := cmd.CombinedOutput()

	if err != nil {
		return
	}

	fmt.Println(string(output))
}
