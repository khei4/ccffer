package info

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type Funcs struct {
	GenericFuncs    []*ast.FuncDecl // 型パラメーターを持つ(引数を持つかはわからない)
	NonGenericFuncs []*ast.FuncDecl // 型パラメーターを持つが引数は持たない
	NoParamFuncs    []*ast.FuncDecl // 型パラメーターも引数も持たない
}

// type Structs struct {
// 	NonGenericStructs
// 	GenericStructs
// }
// type Interfaces struct {
// 	NonGenericInterfaces
// 	GenericInterfaces
// }

type Info struct {
	functions Funcs
	// structs Structs
	// Interfaces Interfaces
}

// printFunc prints function's info
func printFuncDecl(fd *ast.FuncDecl) {
	var name string = fd.Name.Name
	fmt.Println("name=", name)
	var ft ast.FuncType = *fd.Type
	if ft.TypeParams != nil {
		fmt.Println("TypeParams=", *ft.TypeParams)
	}
	if ft.Params != nil {
		fmt.Println("argTypes=", *ft.Params)
	}
	if ft.Results != nil {
		fmt.Println("resTypes=", *ft.Results)
	}
}

func PrintFunctions(info *Info) {
	funcs := info.functions
	for _, funcclass := range [][]*ast.FuncDecl{funcs.GenericFuncs, funcs.NonGenericFuncs, funcs.NoParamFuncs} {
		for _, gf := range funcclass {
			printFuncDecl(gf)
		}
	}
}

func GetInfoFromFiles(filename string) (*Info, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("cannot open %s", filename)
	}

	info := &Info{}

	for _, dec := range f.Decls {
		switch dec.(type) {
		case *ast.FuncDecl:
			fd := dec.(*ast.FuncDecl)

			if fd.Type.TypeParams != nil {
				info.functions.GenericFuncs = append(info.functions.NoParamFuncs, fd)
			} else if fd.Type.Params != nil {
				info.functions.NonGenericFuncs = append(info.functions.NoParamFuncs, fd)
			} else {
				info.functions.NoParamFuncs = append(info.functions.NoParamFuncs, fd)
			}
		default:
			// TODO: get structs and interfaces
		}
	}
	return info, nil
}
