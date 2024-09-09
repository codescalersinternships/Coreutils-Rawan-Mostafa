package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	internal "github.com/codescalersinternships/Coreutils-Rawan-Mostafa/internal"
)

func main() {

	var isLines bool
	var isWords bool
	var isChars bool

	flag.BoolVar(&isLines, "l", false, "prints the lines count only")
	flag.BoolVar(&isWords, "w", false, "prints the words count only")
	flag.BoolVar(&isChars, "c", false, "prints the characters count only")

	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("No file argument passed to the command")
	}
	file:=flag.Args()[0]
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error in reading the file")
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	if !isLines && !isWords && !isChars {
		isLines = true
		isWords = true
		isChars = true
	}

	err = internal.Wc(scanner, isLines, isWords, isChars)
	if err != nil {
		log.Fatal("Error in scanning std input")
	}
	fmt.Println(file)
}
