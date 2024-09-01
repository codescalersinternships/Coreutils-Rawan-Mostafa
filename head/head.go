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
	if n > len(dataArray) {
		n = len(dataArray)
	}
	for i := 0; i < n; i++ {
		fmt.Println(dataArray[i])
	}
}
