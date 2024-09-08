package internal

import (
	"fmt"
	"strings"
)

func Yes(inputs []string) {
	if len(inputs) > 0 {
		input := strings.Join(inputs, " ")
		for {
			fmt.Println(input)
		}
	} else {
		for {
			fmt.Println("y")
		}
	}
}
