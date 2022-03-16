package info_test

import (
	"ccffer/info"
	"testing"
)

func TestGetInfo(t *testing.T) {
	// TODO: テストデータ構造体を作る
	filename := "./.test.go"
	srcinfo, err := info.GetInfoFromFiles(filename)
	if err != nil {
		t.Fatalf("file named %s cannot be open", filename)
	}
	genfs := srcinfo.Functions.GenericFuncs
	nongenfs := srcinfo.Functions.NonGenericFuncs
	noparamfs := srcinfo.Functions.NoParamFuncs
	if len(genfs) != 0 {
		t.Fatal("wrong number: Generic Functions")
	}
	if len(nongenfs) != 1 {
		t.Fatalf("wrong number: %d Non-Generic Functions", len(nongenfs))
	}
	if len(noparamfs) != 3 {
		t.Fatal("wrong number: No-Param Functions")
	}
}
