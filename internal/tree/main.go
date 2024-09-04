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
func formattedPrint(directoryPath string, depth int) {
	fmt.Print("|")

	for i := 0; i < depth; i++ {
		fmt.Print("──")
	}
	fmt.Println(directoryPath)
}
func tree(directoryPath string, depth int, maxDepth int) {

	items, err := os.ReadDir(directoryPath)
	check(err)
	pathParts := strings.Split(directoryPath, "/")
	formattedPrint(pathParts[len(pathParts)-1], depth)

	for _, item := range items {
		if item.IsDir() && depth+1 <= maxDepth {
			tree(directoryPath+"/"+item.Name(), depth+1, maxDepth)
		} else if !item.IsDir() && depth+1 <= maxDepth {
			formattedPrint(item.Name(), depth+1)
		} else if depth+1 > maxDepth {
			return
		}
	}
}

func main() {
	var maxDepth int
	flag.IntVar(&maxDepth, "L", 1, "Print files and directories up to 'L' levels of depth")
	flag.Parse()
	if len(flag.Args()) == 0 {
		currDir, err := os.Getwd()
		check(err)
		tree(currDir, 0, maxDepth)
	} else {
		tree(flag.Args()[0], 0, maxDepth)
	}
}
