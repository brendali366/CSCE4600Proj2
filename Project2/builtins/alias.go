package builtins

import (
	"fmt"
	"os/exec"
)

func AliasCommand(args ...string) error {
	cmd := exec.Command(args[0], args[1])
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error: %v\n%s", err, output)
	}
	fmt.Println(string(output))
	return nil
}
