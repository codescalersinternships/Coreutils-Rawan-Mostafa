package internal

import (
	"bufio"
	"fmt"
	"log"
)

func Head(scanner *bufio.Scanner, n int) {

	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
		if count <= n {
			fmt.Printf("%s \n", string(scanner.Bytes()))
		}
	}
	if scanner.Err() != nil {
		log.Fatal("Error in scanning std input")
	}
}
