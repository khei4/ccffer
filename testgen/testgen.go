package testgen

import "html/template"

type TypeVal struct {
	Type string
	Val  string // interface{} // ここでGenericsにすると, Loopがきつい?
}

type GenFunc struct {
	FName    string
	TypeVals []*TypeVal
}

type TemplData struct {
	Mod      string
	GenFuncs []*GenFunc
}

// できれば, 定義もここで展開したい?
var TestTmpl = template.Must(template.New("gen_test.go").Parse(`
{{$mod := .Mod}}
package {{$mod}}_test

import (
	"testing"
	"{{$mod}}"
)
{{range $gf := .GenFuncs}}
func Test{{$gf.FName}}(t *testing.T) {
	{{range $tv := .TypeVals}}
		{{$mod}}.{{$gf.FName}}[{{$tv.Type}}]({{$tv.Val}})
	{{end}}
}
{{end}}
`))

func GenTests() {}
