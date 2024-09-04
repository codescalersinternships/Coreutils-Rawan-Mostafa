package internal

import (
	"bufio"
	"fmt"
	"log"
)

func Wc(scanner *bufio.Scanner, isLines bool, isWords bool, isChars bool) {

	var linesCount int
	var wordsCount int
	var charsCount int

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		linesCount++
		for _, char := range scanner.Text() {
			charsCount++
			if char == ' ' || char == '\n' || char == '\t' {
				wordsCount++
			}
			if scanner.Err() != nil {
				log.Fatal("Error in scanning std input")
			}
		}
		wordsCount += 1
		charsCount += 1
	}

	if isLines {
		fmt.Printf("%d ", linesCount)
	}
	if isWords {
		fmt.Printf("%d ", wordsCount)
	}
	if isChars {
		fmt.Printf("%d ", charsCount)
	}
	fmt.Print("\n")
}
