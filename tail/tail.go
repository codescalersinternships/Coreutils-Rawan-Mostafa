package main

import (
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
	data, err := os.ReadFile(flag.Args()[0])
	check(err)

	dataArray := strings.Split(string(data), "\n")
	if n >= len(dataArray)-1 {
		for i := 0; i < len(dataArray)-1; i++ {
			fmt.Println(dataArray[i])
		}
	} else {
		for i := len(dataArray) - 1 - n; i < len(dataArray)-1; i++ {
			fmt.Println(dataArray[i])
		}
	}
}
