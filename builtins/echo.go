package builtins

import (
	"fmt"
	//"os"
)

/*
echo in CSH(the C shell) as described in https://docstore.mik.ua/orelly/linux/lnut/ch08_09.htm
echo [-n] string
Write string to standard output; if -n is specified, the output is not terminated by a newline.
... the C shell's echo doesn't support escape characters.
*/
func EchoCSH(args ...string) error {
	// checking if -n flag is raised
	if args[0] == "-n" { //if so, removing before proceeding with echoing WITHOUT newline
		args = args[1:]
		fmt.Println(args)
	} else { //else can echo as is WITH newline
		fmt.Println(args)
		fmt.Println()
	}
	return nil
}
