package main

import (
	"flag"
	"fmt"
)

func main() {

	noNewLine := flag.Bool("n", false, "defines whether or not to add a new line")
	flag.Parse()
	for _, arg := range flag.Args() {
		fmt.Printf("%s ", arg)
	}
	if !*noNewLine {
		fmt.Print("\n")
	}
}
