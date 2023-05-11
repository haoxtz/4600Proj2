package builtins

import (
	"fmt"
	"os"
)

func echoCSH(args []string) {
	arg := os.Args[1:]
	fmt.Println(arg)
}
