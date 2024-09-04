package internal

import (
	"bufio"
	"fmt"
	"log"
)

func Cat(scanner *bufio.Scanner, numbering bool) {

	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		if numbering {
			count++
			fmt.Printf("	%d  %s\n", count, string(scanner.Bytes()))
		} else {
			fmt.Printf("%s\n", string(scanner.Bytes()))
		}
		if scanner.Err() != nil {
			log.Fatal("Error in scanning std input")
		}
	}
}
