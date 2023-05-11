package builtins

import (
	"fmt"
	//"os"
)

func EchoCSH(args ...string) error {
	/*
		echo in CSH(the C shell) as described in https://docstore.mik.ua/orelly/linux/lnut/ch08_09.htm
		echo [-n] string
		Write string to standard output; if -n is specified, the output is not terminated by a newline.
		... the C shell's echo doesn't support escape characters.
	*/
	if args[0] == "-n" {
		args = args[1:]
		fmt.Println(args)
	} else {
		fmt.Println(args)
		fmt.Println()
	}
	return nil
}
