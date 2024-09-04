package internal

import (
	"fmt"
)

func Echo(noNewLine bool, args []string) {

	for _, arg := range args {
		fmt.Printf("%s ", arg)
	}
	if !noNewLine {
		fmt.Print("\n")
	}
}
