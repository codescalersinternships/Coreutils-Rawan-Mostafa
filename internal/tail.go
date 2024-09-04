package internal

import (
	"bufio"
	"fmt"
	"log"
)

func Tail(scanner *bufio.Scanner, n int) {

	scanner.Split(bufio.ScanLines)
	count := 0
	lines := make([]string, 0)
	for scanner.Scan() {
		count++
		lines = append(lines, string(scanner.Bytes()))
	}
	if scanner.Err() != nil {
		log.Fatal("Error in scanning std input")
	}
	if n >= len(lines) {
		n = len(lines)
	}
	for i := len(lines) - n; i < len(lines); i++ {
		fmt.Println(lines[i])
	}

}
