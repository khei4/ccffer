package testgen_test

import (
	"bytes"
	"ccffer/testgen"
	"fmt"
	"testing"
)

func TestGenTests(t *testing.T) {
	appint := &testgen.App{[]testgen.Type{"int"}, []testgen.Val{"0"}}
	appstring := &testgen.App{[]testgen.Type{"string"}, []testgen.Val{`""`}}
	gf := testgen.GenFunc{FName: "F", Apps: []*testgen.App{appint, appstring}}
	td := testgen.TemplData{
		PkgName:  "main",
		GenFuncs: []*testgen.GenFunc{&gf},
	}

	var testsrc bytes.Buffer
	testgen.TestTmpl.Execute(&testsrc, td)
	fmt.Print(testsrc.String())
}
