package main

import (
	"flag"
	"log"
	"os"

	internal "github.com/codescalersinternships/Coreutils-Rawan-Mostafa/internal"
)

func main() {
	var maxDepth int
	flag.IntVar(&maxDepth, "L", 1, "Print files and directories up to 'L' levels of depth")
	flag.Parse()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error reaching the current directory")
	}
	if len(flag.Args()) != 0 {
		dir = flag.Args()[0]
	}
	internal.Tree(dir, 0, maxDepth)
}
