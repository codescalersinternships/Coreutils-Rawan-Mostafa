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

	numbering := flag.Bool("n", false, "defines whether or not the lines will be ordered")
	flag.Parse()
	if len(flag.Args()) == 0 || flag.Args()[0] == "-" {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			text := scanner.Text()
			if text == "exit" {
				break
			}
			fmt.Println(text)
			check(scanner.Err())
		}
	} else {

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
			check(scanner.Err())
		}
	}
}
