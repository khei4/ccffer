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

func classifyFuncDecl(dec ast.Decl, info *Info) {
	fd := dec.(*ast.FuncDecl)

	if fd.Type.TypeParams != nil {
		info.functions.GenericFuncs = append(info.functions.GenericFuncs, fd)
	} else if fd.Type.Params != nil {
		info.functions.NonGenericFuncs = append(info.functions.NonGenericFuncs, fd)
	} else {
		info.functions.NoParamFuncs = append(info.functions.NoParamFuncs, fd)
	}
}

// TODO:
func classifyGenDecl(dec ast.Decl, info *Info) {
	gd := dec.(*ast.GenDecl)
	// Specsが一つ以上ある時とは???
	// →type( ...) が複数あるときっぽい
	for _, spec := range gd.Specs {
		switch spec.(type) {
		case *ast.TypeSpec:
			ts := spec.(*ast.TypeSpec)
			// Genericな変数, 構造体, Interfaceそれぞれあって
			// このレベルでパラメータがある
			// パラメータは別にしてもっとかん?
			// 一旦Genericな構造体, InterfaceTypeは扱わないことに
			if ts.TypeParams != nil {
			}
			switch ts.Type.(type) {
			case *ast.StructType:
			case *ast.InterfaceType:
			default:
			}
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
			classifyFuncDecl(dec, info)
		case *ast.GenDecl:
			// TODO: get structs and interfaces
			classifyGenDecl(dec, info)
		default:
		}
	}
	return info, nil
}
