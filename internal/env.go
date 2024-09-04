package internal

import (
	"os"
)

func Env() {
	for _, envVar := range os.Environ() {
		println(envVar)
	}
}
