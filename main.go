package main

import (
	"fmt"
	"os"

	"github.com/khei4/ccffer/info"
	"github.com/khei4/ccffer/testgen"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "パターンを指定してください")
		return
	}
	// 与えたパスの下を再帰的にとりたい
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
	fp, err := os.Create(os.Args[1] + "/" + os.Args[1] + "_test.go")
	fp.WriteString(test)
	fmt.Print(test)
}
