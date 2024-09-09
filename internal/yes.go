package internal

import (
	"fmt"
	"strings"
)

func Yes(inputs []string) {
	output := "y"
	if len(inputs) > 0 {
		output = strings.Join(inputs, " ")
	}
	for {
		fmt.Println(output)
	}
}
