package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	internal "github.com/codescalersinternships/Coreutils-Rawan-Mostafa/internal"
)

func main() {

	var numbering bool
	flag.BoolVar(&numbering, "n", false, "defines whether or not the lines will be ordered")
	flag.Parse()

	var scanner *bufio.Scanner

	if len(flag.Args()) == 0 || flag.Args()[0] == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		data, err := os.ReadFile(flag.Args()[0])
		if err != nil {
			log.Fatal("Error in reading the file")
		}

		scanner = bufio.NewScanner(strings.NewReader(string(data)))
	}
	internal.Cat(scanner, numbering)
}
