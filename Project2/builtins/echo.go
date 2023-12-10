package builtins

import (
	"fmt"
	"os/exec"
)

// takes in each input after echo and outputs it
func EchoCommands(args ...string) {
	for i := 1; i < len(args); i++ {
		cmd := exec.Command(args[i])
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//this is where it is output-ed
		fmt.Println(string(stdout))
	}
}
