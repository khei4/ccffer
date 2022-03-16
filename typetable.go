// このType Fuzzingで試す型一覧
// https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
package main

var TypeVals map[string][]string = map[string][]string{
	"bool":   {"true", "false"},
	"string": {""},
	"int":    {"0", "-(1 << 31) - 1", "1 << 31"},
	"int64":  {"0", "-(1 << 63) - 1", "1 << 63"},
	"uint":   {"0", "^uint(0)"},
	// TODO: import math library by formatting test
	// "float32": []string{"0", "math.NaN()", "math.Inf(0)", "math.Inf(-1)"},
}
