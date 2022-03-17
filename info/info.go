package info

import (
	"ccffer/model"
	"ccffer/typetable"
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/packages"
)

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

func argall(argcands [][]model.Val) [][]model.Val {
	args := make([][]model.Val, 1)
	for _, l := range argcands {
		newargs := make([][]model.Val, 0)
		for _, arg := range args {
			for _, v := range l {
				newargs = append(newargs, append(arg, v))
			}
		}
		args = newargs
	}
	return args
}

// 考えられる入力と型を列挙して入れていく
func getFuzzAppsFuncDecl(pkg *packages.Package, fd *ast.FuncDecl, tmplData *model.TemplData) {
	sig, _ := pkg.TypesInfo.TypeOf(fd.Name).(*types.Signature)
	if sig == nil {
		return
	}
	var name string = fd.Name.Name
	gf := &model.GenFunc{FName: name}
	// 引数が1個以上かつ, Generic
	if fd.Type.TypeParams != nil && len(fd.Type.Params.List) != 0 {

		// 1. それぞれの型パラメーターに入る型を列挙する.
		//		1. 型パラメーターに入れられる型の候補を列挙する.
		//		2. 型制約を満たすかどうか判定する
		// typpnums := len(fd.Type.TypeParams.List)
		// typecands := typecandsGen()
		if _, err := types.Instantiate(nil, sig, []types.Type{types.Typ[types.Int]}, true); err != nil {
			fmt.Println("can't be instantiated by int ")
		}
		if _, err := types.Instantiate(nil, sig, []types.Type{types.Typ[types.Bool]}, true); err != nil {
			fmt.Println("can't be instantiated by bool")
		}
		if _, err := types.Instantiate(nil, sig, []types.Type{types.Typ[types.String]}, true); err != nil {
			fmt.Println("can't be instantiated by string")
		}

		// 2. 型を固定して, 下と同じ操作
		// types.Instantiate()
		// その型を固定してLoop
		// Generic
		// types.TypeParam
		// for typ, vals := range typetable.TypeVals {
		// }
	} else if len(fd.Type.Params.List) != 0 {
		argcands := make([][]model.Val, len(fd.Type.Params.List))
		for i, field := range fd.Type.Params.List {
			switch typ := pkg.TypesInfo.TypeOf(field.Type).(type) {
			case *types.Basic: // basic typeを入れる
				// 本当は全通り入れたい
				argcands[i] = typetable.TypeVals[typ]
			case *types.Pointer: // nil をいれる
				argcands[i] = []model.Val{"nil"}
			case *types.Interface: //  nil + TODO
				argcands[i] = []model.Val{"nil"}
			case *types.TypeParam:
				panic("broken")
			}
		}
		for _, args := range argall(argcands) {
			gf.Apps = append(gf.Apps, &model.App{TypeInstances: []model.Type{}, Args: args})
		}

	} else {
	}
	tmplData.GenFuncs = append(tmplData.GenFuncs, gf)

}

func GetTemplDataFromPackages(pkgnames []string) (*model.TemplData, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(cfg, pkgnames...)
	if err != nil {
		return nil, err
	}
	// TODO:単一packageを仮定しているので拡張する.
	templData := &model.TemplData{PkgName: pkgs[0].Name}
	for _, pkg := range pkgs {
		// templData := &model.TemplData{PkgName: pkg.Name}
		for _, f := range pkg.Syntax {
			// ast.Print(pkg.Fset, f)
			for _, dec := range f.Decls {
				switch dec := dec.(type) {
				case *ast.FuncDecl:
					getFuzzAppsFuncDecl(pkg, dec, templData)
				case *ast.GenDecl:
					// TODO: get structs and interfaces
				default:
				}
			}
		}
	}

	return templData, nil
}
