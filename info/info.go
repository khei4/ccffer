package info

import (
	"ccffer/model"
	"ccffer/typetable"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/packages"
)

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

func typecandsGen(fd *ast.FuncDecl) [][]types.Type {
	num := len(fd.Type.TypeParams.List)
	// res := make([][]types.Type, 0, candsnum)
	argtypes := make([][]types.Type, 1)
	for i := 0; i < num; i++ {
		// allocate
		newargtypes := make([][]types.Type, 0, len(argtypes)*len(typetable.TypeVals))
		for _, argtype := range argtypes {
			for typ, _ := range typetable.TypeVals {
				newargtypes = append(newargtypes, append(argtype, typ))
			}
		}
		argtypes = newargtypes
	}
	return argtypes
}

func types2strings(typs []types.Type) []model.Type {
	res := make([]model.Type, len(typs))
	for i, t := range typs {
		res[i] = t.String()
	}
	return res
}

// 考えられる入力と型を列挙して入れていく
func getFuzzAppsFuncDecl(pkg *packages.Package, fd *ast.FuncDecl, tmplData *model.TemplData) {
	sig, _ := pkg.TypesInfo.TypeOf(fd.Name).(*types.Signature)
	if sig == nil {
		return
	}
	var name string = fd.Name.Name
	gf := &model.GenFunc{FName: name}
	if fd.Type.TypeParams != nil && len(fd.Type.Params.List) != 0 {
		typecands := typecandsGen(fd)
		// TODO: 長さごとにCashに乗せるなりしないとやばそう
		allowedcands := make([][]types.Type, 0)

		for _, cand := range typecands {
			if _, err := types.Instantiate(nil, sig, cand, true); err != nil {
				// fmt.Fprintf(os.Stderr, "%v can't be instantiated by %v\n", name, cand)
				continue
			}
			allowedcands = append(allowedcands, cand)
		}
		for _, typecands := range allowedcands {
			cur := 0
			argcands := make([][]model.Val, len(fd.Type.Params.List))
			for i, field := range fd.Type.Params.List {
				switch typ := pkg.TypesInfo.TypeOf(field.Type).(type) {
				case *types.Basic: // basic typeを入れる
					argcands[i] = typetable.TypeVals[typ]
				case *types.Pointer: // nil をいれる
					argcands[i] = []model.Val{"nil"}
				case *types.Interface: //  nil + TODO
					argcands[i] = []model.Val{"nil"}
				case *types.TypeParam:
					argcands[i] = typetable.TypeVals[typecands[cur]]
				}
			}
			typeInstances := types2strings(typecands)
			for _, args := range argall(argcands) {
				gf.Apps = append(gf.Apps, &model.App{
					TypeInstances: typeInstances,
					Args:          args,
				})
			}
		}

	} else if len(fd.Type.Params.List) != 0 {
		argcands := make([][]model.Val, len(fd.Type.Params.List))
		for i, field := range fd.Type.Params.List {
			switch typ := pkg.TypesInfo.TypeOf(field.Type).Underlying().(type) {
			case *types.Basic:
				argcands[i] = typetable.TypeVals[typ]
			case *types.Slice:
				elemType := typ.Elem()
				empty := "[]" + elemType.String() + "{}"
				argcands[i] = []model.Val{"nil", empty}
			case *types.Pointer:
				argcands[i] = []model.Val{"nil"}
			case *types.Interface:
				argcands[i] = []model.Val{"nil"}
			case *types.Struct:
				zeroVal := typ.String() + "{}"
				argcands[i] = []model.Val{"nil", zeroVal}
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
		Mode: packages.NeedModule | packages.NeedName | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(cfg, pkgnames...)
	if err != nil {
		return nil, err
	}
	// TODO:単一packageを仮定しているので拡張する.
	templData := &model.TemplData{PkgName: pkgs[0].Name, PkgPath: pkgs[0].PkgPath}
	for _, pkg := range pkgs {
		// templData := &model.TemplData{PkgName: pkg.Name}
		for _, f := range pkg.Syntax {
			// ast.Print(pkg.Fset, f)
			for _, dec := range f.Decls {
				switch dec := dec.(type) {
				case *ast.FuncDecl:
					getFuzzAppsFuncDecl(pkg, dec, templData)
				case *ast.GenDecl:
					// TODO: get structs and interfaces and add them to table

				default:
				}
			}
		}
	}

	return templData, nil
}
