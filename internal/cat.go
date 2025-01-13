package internal

import (
	"bufio"
	"fmt"
)

func Cat(scanner *bufio.Scanner, numbering bool) error {

	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		if numbering {
			count++
			fmt.Printf("	%d  %s\n", count, string(scanner.Bytes()))
		} else {
			fmt.Printf("%s\n", string(scanner.Bytes()))
		}
	}
	return scanner.Err()
}
