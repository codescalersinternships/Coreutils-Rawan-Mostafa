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

	var n int

	flag.IntVar(&n, "n", 10, "defines the number fo lines to be printed")
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("No file argument passed to the command")
	}
	data, err := os.ReadFile(flag.Args()[0])
	if err != nil {
		log.Fatal("Error in reading the file")
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	internal.Tail(scanner, n)
}
