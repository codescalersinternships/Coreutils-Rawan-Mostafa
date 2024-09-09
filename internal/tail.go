package internal

import (
	"bufio"
	"fmt"
)

func Tail(scanner *bufio.Scanner, n int) error {

	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, string(scanner.Bytes()))
	}

	if n >= len(lines) {
		n = len(lines)
	}
	for i := len(lines) - n; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
	return scanner.Err()
}
