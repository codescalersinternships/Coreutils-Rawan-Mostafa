package internal

import (
	"fmt"
	"strings"
)

func Echo(noNewLine bool, args []string) {

	argString := strings.Join(args, " ")
	fmt.Print(argString)
	if !noNewLine {
		fmt.Print("\n")
	}
}
