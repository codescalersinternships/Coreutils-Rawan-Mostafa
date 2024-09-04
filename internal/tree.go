package internal

import (
	"fmt"
	"log"
	"os"
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
	pathParts := strings.Split(directoryPath, "/")
	formattedPrint(pathParts[len(pathParts)-1], depth)

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
