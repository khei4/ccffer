package testgen_test

import (
	"bytes"
	"ccffer/model"
	"ccffer/testgen"
	"testing"
)

func TestGenTests(t *testing.T) {
	appint := &model.App{TypeInstances: []model.Type{"int"}, Args: []model.Val{"0"}}
	appstring := &model.App{TypeInstances: []model.Type{"string"}, Args: []model.Val{`""`}}
	gf := model.GenFunc{FName: "F", Apps: []*model.App{appint, appstring}}
	td := model.TemplData{
		PkgName:  "main",
		GenFuncs: []*model.GenFunc{&gf},
	}

	var testsrc bytes.Buffer
	testgen.TestTmpl.Execute(&testsrc, td)
	expectedraw := `package main_test
import (
        "testing"
        "main"
)
func TestF(t *testing.T) {
	main.F[int,](0,)
	main.F[string,]("",)
}
`

	expected, _ := testgen.FormatAndImport([]byte(expectedraw))
	if got, err := testgen.GenTests(&td); got != expected || err != nil {
		t.Fatalf("%v\n is not \n%v", got, expected)
	}

}
