package internal

import (
	"bufio"
	"fmt"
)

func Head(scanner *bufio.Scanner, n int) error {

	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
		if count <= n {
			fmt.Printf("%s \n", string(scanner.Bytes()))
		}
	}
	return scanner.Err()
}
