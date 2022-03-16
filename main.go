package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	source := `
	package main
	import "fmt"
	func main2() (){
		fmt.Print("hello")
	}
	func main3() (){
		fmt.Print("hola")
	}
	func main4() (){
		fmt.Print("bon dia")
	}
	`
	filename := "hoge.go"
	f, err := parser.ParseFile(fset, filename, source, parser.ParseComments)

	for _, dec := range f.Decls {
		switch dec.(type) {
		case *ast.FuncDecl:
			fd := dec.(*ast.FuncDecl)
			var name string = fd.Name.Name
			fmt.Println("name=", name)
			var ft ast.FuncType = *fd.Type

			// var typeParams ast.FieldList = *ft.TypeParams
			var argTypes ast.FieldList = *ft.Params
			var retTypes ast.FieldList = *ft.Results
			// fmt.Print(typeParams)
			fmt.Println("argTypes=", argTypes)
			fmt.Print(retTypes)
		default:
		}
	}
	_ = err
}
