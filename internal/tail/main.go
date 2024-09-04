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

	var n int

	flag.IntVar(&n, "n", 10, "defines the number fo lines to be printed")
	flag.Parse()
	if len(flag.Args()) == 0 {
		panic("No file argument passed to the command")
	}
	data, err := os.ReadFile(flag.Args()[0])
	check(err)

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	scanner.Split(bufio.ScanLines)
	count := 0
	lines := make([]string, 0)
	for scanner.Scan() {
		count++
		lines = append(lines, string(scanner.Bytes()))
	}
	check(scanner.Err())
	if n >= len(lines) {
		for i := 0; i < len(lines); i++ {
			fmt.Println(lines[i])
		}
	} else {
		for i := len(lines) - n; i < len(lines); i++ {
			fmt.Println(lines[i])
		}

	}
}
