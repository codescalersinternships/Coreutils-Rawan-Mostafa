package main

import (
	"flag"

	internal "github.com/codescalersinternships/Coreutils-Rawan-Mostafa/internal"
)

func main() {
	var noNewLine bool

	flag.BoolVar(&noNewLine, "n", false, "defines whether or not to add a new line")
	flag.Parse()

	internal.Echo(noNewLine, flag.Args())
}
