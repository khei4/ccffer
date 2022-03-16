package testgen

import (
	"html/template"
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
	Mod      string
	GenFuncs []*GenFunc
}

// 複数変数, 複数パラメーター
// 愚直なfor文だと失敗する
var TestTmpl = template.Must(template.New("gen_test.go").Funcs(template.FuncMap{
	"createCS": func(ss []string) string {
		var res = make([]byte, 0, 100)
		for _, s := range ss {
			res = append(res, s...)
			res = append(res, ',')
		}
		return string(res)
	}}).Parse(`
{{$mod := .Mod}}
package {{$mod}}_test

import (
	"testing"
	"{{$mod}}"
)
{{range $gf := .GenFuncs}}
func Test{{$gf.FName}}(t *testing.T) {
	{{range $app := $gf.Apps}}
		{{$mod}}.{{$gf.FName}}[{{ createCS $app.TypeInstances}}]({{createCS $app.Args}})
	{{end}}
}
{{end}}
`))

func GenTests() {

}
