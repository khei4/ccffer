package testgen

import (
	"text/template"
)

type Val = string
type Type = string

type App struct {
	TypeInstances []Type // 型パラメーターの値
	Args          []Val  // 引数にわたす値
}

type GenFunc struct {
	FName string
	Apps  []*App
}

type TemplData struct {
	PkgName  string
	GenFuncs []*GenFunc
}

var TestTmpl = template.Must(template.New("gen_test.go").Funcs(template.FuncMap{
	"createCS": func(ss []string) string {
		var res = make([]byte, 0, 100)
		for _, s := range ss {
			res = append(res, s...)
			res = append(res, ',')
		}
		return string(res)
	}}).Parse(`
	{{$pkg := .PkgName}}
	package {{$pkg}}_test

	import (
		"testing"
		"{{$pkg}}"
	)
	{{range $gf := .GenFuncs}}
	func Test{{$gf.FName}}(t *testing.T) {
		{{range $app := $gf.Apps}}
			{{$pkg}}.{{$gf.FName}}[{{ createCS $app.TypeInstances}}]({{createCS $app.Args}})
		{{end}}
	}
	{{end}}
`))

func GenTests() {

}
