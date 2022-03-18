// このType Fuzzingで試す型一覧
// https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
package typetable

import "go/types"

// TODO: Empty Interfaceの型がほしい
var TypeVals map[types.Type][]string = map[types.Type][]string{
	types.Typ[types.Bool]:                  {"true", "false"},
	types.Typ[types.String]:                {"\"\"", "\"hogehogehogehoge\""},
	types.Typ[types.Int]:                   {"0", "-2147483648", "2147483647"},
	types.Typ[types.Int64]:                 {"0", "-9223372036854775808", "9223372036854775807"},
	types.Typ[types.Uint]:                  {"0", "18446744073709551615"},
	&types.Struct{}:                        {"struct{}{}"},
	types.Typ[types.Float32]:               {"0", "float32(math.NaN())", "float32(math.Inf(0))", "float32(math.Inf(-1))"},
	types.Typ[types.Float64]:               {"0", "math.NaN()", "math.Inf(0)", "math.Inf(-1)"},
	types.NewPointer(types.Typ[types.Int]): {"&0", "nil"},
	types.NewSlice(types.Typ[types.Int]):   {"[]int{}", "nil"},
	// &types.Pointer{}:        {"nil"},
	// &types.Slice{}:          {"[]", "nil"},
}
