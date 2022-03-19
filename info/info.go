package info

import (
	"go/ast"
	"go/types"
	"strings"
	"unicode"

	"github.com/khei4/ccffer/model"
	"github.com/khei4/ccffer/typetable"

	"golang.org/x/tools/go/packages"
)

// calculate combinations from candidates repeated rep times. For examples
//		In: [["1", "2"], ["", "hoge"]], 1
// 		Out: [["1", ""], ["2", ""], ["1", "hoge"], [2, "hoge"]]
// 		In: typetable.TypeVals, 2
// 		Out: [[types.Typ[types.Bool], types.Typ[types.Bool]],
// 				 ...
//			[types.NewPointer(types.Typ[types.Int]), types.NewSlice(types.Typ[types.Int])],
// 			[types.NewSlice(types.Typ[types.Int]),types.NewSlice(types.Typ[types.Int])]]
func constructCombinations[T any](candidates [][]T, rep int) [][]T {
	combs := make([][]T, 1)
	for i := 0; i < rep; i++ {
		for _, l := range candidates {
			newcombs := make([][]T, 0, len(combs)*len(l))
			for _, arg := range combs {
				for _, v := range l {
					newcombs = append(newcombs, append(arg, v))
				}
			}
			combs = newcombs
		}
	}
	return combs
}

// getTypeParamNum get a number of typeparameter for given fd *ast.FuncDecl
func getTypeParamNum(fd *ast.FuncDecl) int {
	num := 0
	for _, l := range fd.Type.TypeParams.List {
		num += len(l.Names)
	}
	return num
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
	var name string = fd.Name.Name
	// ここはよりよいPublic判定方法があるかな?
	if name == "main" || unicode.IsLower([]rune(name)[0]) {
		return
	}
	gf := &model.GenFunc{FName: name}
	sig, _ := pkg.TypesInfo.TypeOf(fd.Name).(*types.Signature)
	if sig == nil {
		return
	}

	if fd.Type.TypeParams != nil && len(fd.Type.Params.List) != 0 {
		num := getTypeParamNum(fd)
		typecands := constructCombinations([][]types.Type{typetable.Types}, num)
		// TODO: 長さごとにCashに乗せるなりしないと引数が多い時にやばそう
		allowedcands := make([][]types.Type, 0)

		for _, cand := range typecands {
			if _, err := types.Instantiate(nil, sig, cand, true); err != nil {
				continue
			}
			allowedcands = append(allowedcands, cand)
		}
		for _, typecands := range allowedcands {
			cur := 0
			argcands := make([][]model.Val, len(fd.Type.Params.List))
			for i, field := range fd.Type.Params.List {
				switch typ := pkg.TypesInfo.TypeOf(field.Type).(type) {
				case *types.Basic:
					argcands[i] = typetable.TypeVals[typ]
				case *types.Pointer:
					argcands[i] = []model.Val{"nil"}
				case *types.Interface:
					argcands[i] = []model.Val{"nil"}
				case *types.TypeParam:
					argcands[i] = typetable.TypeVals[typecands[cur]]
					cur += 1
				}
			}
			typeInstances := types2strings(typecands)
			for _, args := range constructCombinations(argcands, 1) {
				gf.Apps = append(gf.Apps, &model.App{
					TypeInstances: typeInstances,
					Args:          args,
				})
			}
		}

	} else if len(fd.Type.Params.List) != 0 {
		argcands := make([][]model.Val, len(fd.Type.Params.List))
		for i, field := range fd.Type.Params.List {
			switch typ := pkg.TypesInfo.TypeOf(field.Type).(type) {
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
				argcands[i] = []model.Val{zeroVal}
			case *types.Named:
				path := strings.Split(typ.String(), "/")
				zeroVal := path[len(path)-1] + "{}"
				argcands[i] = []model.Val{zeroVal}
			case *types.TypeParam:
				panic("broken")
			}
		}
		for _, args := range constructCombinations(argcands, 1) {
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
