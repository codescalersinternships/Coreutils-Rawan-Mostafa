package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func formattedPrint(directoryPath string, depth int) {
	fmt.Print("|")
	fmt.Print(strings.Repeat("──", depth))
	fmt.Println(directoryPath)
}
func Tree(directoryPath string, depth int, maxDepth int) {

	items, err := os.ReadDir(directoryPath)
	if err != nil {
		log.Fatalf("Error reading items in: %s", directoryPath)
	}
	lastPathElement := filepath.Base(directoryPath)
	formattedPrint(lastPathElement, depth)

	for _, item := range items {
		if depth+1 > maxDepth {
			return
		} else if item.IsDir() {
			Tree(directoryPath+"/"+item.Name(), depth+1, maxDepth)
		} else {
			formattedPrint(item.Name(), depth+1)
		}
	}
}
