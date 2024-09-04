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
	if len(flag.Args()) == 0 {
		currDir, err := os.Getwd()
		if err != nil {
			log.Fatal("Error reaching the current directory")
		}
		internal.Tree(currDir, 0, maxDepth)
	} else {
		internal.Tree(flag.Args()[0], 0, maxDepth)
	}
}
