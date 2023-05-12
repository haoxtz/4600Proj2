package builtins

import (
	"errors"
	"fmt"
	"os"
	
)


func RemoveFile(args ...string) error {


	if len(args) == 0 {
		return fmt.Errorf("invalid argument count") // if no arguments provided, return an error
	}

	var allErrors []string
	for _, arg := range args {
		err := os.Remove(arg) // tries to remove the file at the given path
		if err != nil {
			allErrors = append(allErrors, err.Error()) // if error occurs, return error message
		}
	}

	if len(allErrors) > 0 {
		return errors.New(fmt.Sprintf("%s", allErrors)) 
	}

	return nil 
}
