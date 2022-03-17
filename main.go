package main

import (
	"fmt"
	"os"

	"ccffer/info"
	"ccffer/testgen"
)

func main() {

	filename := []string{"./.test"}
	templData, err := info.GetTemplDataFromPackages(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open %s", filename)
		return
	}
	test, err := testgen.GenTests(templData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error: %s", err)
	}
	fmt.Print(test)
}
