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

	numbering := flag.Bool("n", false, "defines the number fo lines to be printed")
	flag.Parse()
	data, err := os.ReadFile(flag.Args()[0])
	check(err)

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
		if *numbering {
			fmt.Printf("	%d  %s\n", count, string(scanner.Bytes()))
		} else {
			fmt.Printf("%s\n", string(scanner.Bytes()))
		}
	}
	check(scanner.Err())
}
