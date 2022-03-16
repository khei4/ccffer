package main

import (
	"fmt"
	"os"

	"github.com/khei4/ccffer/info"
)

func genzero[T interface {
	int
}](a, b T) T {
	return a / b
}

func main() {
	filename := "gotest"
	info, err := info.GetInfoFromFiles(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open %s", filename)
	}
	info.PrintFunctions(info)
	fmt.Print(info)
}
