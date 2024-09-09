package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	internal "github.com/codescalersinternships/Coreutils-Rawan-Mostafa/internal"
)

func main() {

	var n int

	flag.IntVar(&n, "n", 10, "defines the number fo lines to be printed")
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("No file argument passed to the command")
	}

	filePath := flag.Args()[0]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error in opening the file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	err = internal.Tail(scanner, n)
	if err != nil {
		log.Fatal("Error in scanning std input")
	}
}
