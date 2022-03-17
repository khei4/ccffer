package info_test

import (
	"ccffer/info"
	"testing"
)

func TestGetInfo(t *testing.T) {
	pkgnames := []string{"../.test"}
	srcinfo, err := info.GetInfoFromPackages(pkgnames)
	if err != nil {
		t.Fatalf("packages cannot be open")
	}
	genfs := srcinfo.Functions.GenericFuncs
	nongenfs := srcinfo.Functions.NonGenericFuncs
	noparamfs := srcinfo.Functions.NoParamFuncs
	if len(genfs) != 1 {
		t.Fatalf("wrong number: %d Generic Functions", len(genfs))
	}
	if len(nongenfs) != 0 {
		t.Fatalf("wrong number: %d Non-Generic Functions", len(nongenfs))
	}
	if len(noparamfs) != 1 {
		t.Fatalf("wrong number: %d No-Param Functions", len(noparamfs))
	}
}
