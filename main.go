package main

import (
	"fmt"
	"os"

	"ccffer/info"
)

func main() {

	filename := []string{"./.test"}
	srcinfo, err := info.GetInfoFromPackages(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open %s", filename)
	}
	info.PrintFunctions(srcinfo)
	fmt.Print(srcinfo)
}
