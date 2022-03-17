package info

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/packages"
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
	Functions Funcs
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
	funcs := info.Functions
	for _, funcclass := range [][]*ast.FuncDecl{funcs.GenericFuncs, funcs.NonGenericFuncs, funcs.NoParamFuncs} {
		for _, gf := range funcclass {
			printFuncDecl(gf)
		}
	}
}

func classifyFuncDecl(pkg *packages.Package, fd *ast.FuncDecl, info *Info) {
	sig, _ := pkg.TypesInfo.TypeOf(fd.Name).(*types.Signature)
	if sig == nil {
		return
	}

	if fd.Type.TypeParams != nil { // TODO: 型引数はあっても, 引数がないのは?
		// for typ, val := range typetable.TypeVals {
		// 	// sigの型パラメーターごとにあてはまるものを列挙
		// 	if
		// }
		info.Functions.GenericFuncs = append(info.Functions.GenericFuncs, fd)
	} else if len(fd.Type.Params.List) != 0 {
		info.Functions.NonGenericFuncs = append(info.Functions.NonGenericFuncs, fd)
	} else {
		info.Functions.NoParamFuncs = append(info.Functions.NoParamFuncs, fd)
	}

}

// TODO:
func classifyGenDecl(gd *ast.GenDecl, info *Info) {
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

func GetInfoFromPackages(pkgnames []string) (*Info, error) {
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(cfg, pkgnames...)
	if err != nil {
		return nil, err
	}
	info := &Info{}
	for _, pkg := range pkgs {
		for _, f := range pkg.Syntax {
			ast.Print(pkg.Fset, f)
			for _, dec := range f.Decls {
				switch dec := dec.(type) {
				case *ast.FuncDecl:
					classifyFuncDecl(pkg, dec, info)
				case *ast.GenDecl:
					// TODO: get structs and interfaces
					classifyGenDecl(dec, info)
				default:
				}
			}
		}
	}

	return info, nil
}
