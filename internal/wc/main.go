package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	linesOnly := flag.Bool("l", false, "prints the lines count only")
	wordsOnly := flag.Bool("w", false, "prints the words count only")
	charsOnly := flag.Bool("c", false, "prints the characters count only")

	flag.Parse()
	if len(flag.Args()) == 0 {
		panic("No file argument passed to the command")
	}
	data, err := os.ReadFile(flag.Args()[0])
	check(err)

	if !*linesOnly && !*wordsOnly && !*charsOnly {
		*linesOnly = true
		*wordsOnly = true
		*charsOnly = true
	}

	if *linesOnly {
		scanner := bufio.NewScanner(strings.NewReader(string(data)))
		scanner.Split(bufio.ScanLines)
		count := 0
		for scanner.Scan() {
			count++
		}
		check(scanner.Err())
		fmt.Printf("%d ", count)
	}
	if *wordsOnly {
		scanner := bufio.NewScanner(strings.NewReader(string(data)))

		scanner.Split(bufio.ScanWords)

		count := 0
		for scanner.Scan() {
			count++
		}
		check(scanner.Err())
		fmt.Printf("%d ", count)
	}
	if *charsOnly {
		scanner := bufio.NewScanner(strings.NewReader(string(data)))
		scanner.Split(bufio.ScanBytes)
		count := 0
		for scanner.Scan() {
			count++
		}
		check(scanner.Err())
		fmt.Printf("%d ", count)
	}

}
