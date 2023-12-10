package builtins

import (
	"fmt"
	"strconv"
)

func RepeatCommand(args ...string) {
	str := args[2]
	num, err := strconv.Atoi(args[0])
	for i := 0; i < num; i++ {
		fmt.Println(str)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
