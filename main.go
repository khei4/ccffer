package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"ccffer/info"
)

func DumpAst(filename string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, f)

}

func main() {
	filename := ".test.go"
	DumpAst(".test.go")
	srcinfo, err := info.GetInfoFromFiles(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open %s", filename)
	}
	info.PrintFunctions(srcinfo)
	fmt.Print(srcinfo)
}
