package testgen_test

import (
	"bytes"
	"ccffer/testgen"
	"fmt"
	"testing"
)

func TestGenTests(t *testing.T) {
	tvint := &testgen.TypeVal{"int", "0"}
	tvstring := &testgen.TypeVal{"string", "\"hoge\""} // TODO: ここの""が&#34;になってしまう
	gf := testgen.GenFunc{FName: "F", TypeVals: []*testgen.TypeVal{tvint, tvstring}}
	td := testgen.TemplData{
		Mod:      "main",
		GenFuncs: []*testgen.GenFunc{&gf},
	}

	var testsrc bytes.Buffer
	testgen.TestTmpl.Execute(&testsrc, td)
	fmt.Print(testsrc.String())
	// t.Log(testsrc.String())
}

// type TypeVal struct {
// 	Type string
// 	Val  string // interface{} // ここでGenericsにすると, Loopがきつい?
// }

// type GenFunc struct {
// 	FName    string
// 	TypeVals []*TypeVal
// }

// type TmplData struct {
// 	Mod      string
// 	GenFuncs []*GenFunc
// }
