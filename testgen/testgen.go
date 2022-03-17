package testgen

import (
	"bytes"
	"ccffer/model"
	"fmt"
	"go/format"
	"text/template"

	"golang.org/x/tools/imports"
)

var TestTmpl = template.Must(template.New("gen_test.go").Funcs(template.FuncMap{
	"createCS": func(ss []string) string {
		var res = make([]byte, 0, 100)
		for _, s := range ss {
			res = append(res, s...)
			res = append(res, ',')
		}
		return string(res)
	}}).Parse(`{{$pkg := .PkgName}}
package {{$pkg}}_test
import (
	"testing"
	"{{$pkg}}"
)
{{range $gf := .GenFuncs}}func Test{{$gf.FName}}(t *testing.T) {
	{{range $app := $gf.Apps}}{{$pkg}}.{{$gf.FName}} {{if lt 0 (len $app.TypeInstances)}} [ {{ createCS $app.TypeInstances}}] {{end}} ({{createCS $app.Args}})
	{{end}}
}{{end}}`))

func FormatAndImport(src []byte) (string, error) {
	srcbytes, err := format.Source(src)
	if err != nil {
		return "", err
	}
	srcbytes, err = imports.Process("", srcbytes, nil)
	if err != nil {
		return "", err
	}
	return string(srcbytes), nil
}

func GenTests(td *model.TemplData) (string, error) {
	var testsrc bytes.Buffer
	TestTmpl.Execute(&testsrc, td)
	fmt.Print(testsrc.String())
	return FormatAndImport(testsrc.Bytes())

}
