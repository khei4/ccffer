// このType Fuzzingで試す型一覧
// https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
package main

import "go/types"

// 値でもったほうがいい?
// Pointer, Interfaceは実際に取ってきてから作らないと厳しそう
var TypeVals map[types.Type][]string = map[types.Type][]string{
	types.Typ[types.Bool]:   {"true", "false"},
	types.Typ[types.String]: {""},
	types.Typ[types.Int]:    {"0", "-2147483648", "2147483647"},
	types.Typ[types.Int64]:  {"0", "-9223372036854775808", "9223372036854775807"},
	types.Typ[types.Uint]:   {"0", "18446744073709551615"},
	// TODO: import math library by formatting test
	// "float32": []string{"0", "math.NaN()", "math.Inf(0)", "math.Inf(-1)"},
}
