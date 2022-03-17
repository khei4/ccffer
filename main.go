package main

import (
	"fmt"
	"os"

	"ccffer/info"
	"ccffer/testgen"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "パターンを指定してください")
		return
	}
	templData, err := info.GetTemplDataFromPackages(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open %s", os.Args[1:])
		return
	}
	test, err := testgen.GenTests(templData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error: %s", err)
		return
	}
	fmt.Print(test)
}
