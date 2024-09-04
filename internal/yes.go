package internal

import (
	"fmt"
	"strings"
)

func Yes(inputs []string) {
	input := strings.Join(inputs, " ")
	for {
		fmt.Println(input)
	}
}
