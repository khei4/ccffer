// Code generated by ccffer; DO NOT EDIT.

package test_test

import (
	"github.com/khei4/ccffer/test"
	"math"
	"testing"
)

func TestInsert(t *testing.T) {
	test.Insert[string, struct{}]("", struct{}{})
	test.Insert[string, int]("", 0)
	test.Insert[string, int]("", -2147483648)
	test.Insert[string, int]("", 2147483647)
	test.Insert[string, int64]("", 0)
	test.Insert[string, int64]("", -9223372036854775808)
	test.Insert[string, int64]("", 9223372036854775807)
	test.Insert[string, uint]("", 0)
	test.Insert[string, uint]("", 18446744073709551615)
	test.Insert[string, float32]("", 0)
	test.Insert[string, float32]("", float32(math.NaN()))
	test.Insert[string, float32]("", float32(math.Inf(0)))
	test.Insert[string, float32]("", float32(math.Inf(-1)))
	test.Insert[string, float64]("", 0)
	test.Insert[string, float64]("", math.NaN())
	test.Insert[string, float64]("", math.Inf(0))
	test.Insert[string, float64]("", math.Inf(-1))
	test.Insert[string, *int]("", nil)
	test.Insert[string, bool]("", true)
	test.Insert[string, bool]("", false)
	test.Insert[string, string]("", "")
	test.Insert[string, []int]("", []int{})
	test.Insert[string, []int]("", nil)
	test.Insert[int, uint](0, 0)
	test.Insert[int, uint](0, 18446744073709551615)
	test.Insert[int, uint](-2147483648, 0)
	test.Insert[int, uint](-2147483648, 18446744073709551615)
	test.Insert[int, uint](2147483647, 0)
	test.Insert[int, uint](2147483647, 18446744073709551615)
	test.Insert[int, float32](0, 0)
	test.Insert[int, float32](0, float32(math.NaN()))
	test.Insert[int, float32](0, float32(math.Inf(0)))
	test.Insert[int, float32](0, float32(math.Inf(-1)))
	test.Insert[int, float32](-2147483648, 0)
	test.Insert[int, float32](-2147483648, float32(math.NaN()))
	test.Insert[int, float32](-2147483648, float32(math.Inf(0)))
	test.Insert[int, float32](-2147483648, float32(math.Inf(-1)))
	test.Insert[int, float32](2147483647, 0)
	test.Insert[int, float32](2147483647, float32(math.NaN()))
	test.Insert[int, float32](2147483647, float32(math.Inf(0)))
	test.Insert[int, float32](2147483647, float32(math.Inf(-1)))
	test.Insert[int, float64](0, 0)
	test.Insert[int, float64](0, math.NaN())
	test.Insert[int, float64](0, math.Inf(0))
	test.Insert[int, float64](0, math.Inf(-1))
	test.Insert[int, float64](-2147483648, 0)
	test.Insert[int, float64](-2147483648, math.NaN())
	test.Insert[int, float64](-2147483648, math.Inf(0))
	test.Insert[int, float64](-2147483648, math.Inf(-1))
	test.Insert[int, float64](2147483647, 0)
	test.Insert[int, float64](2147483647, math.NaN())
	test.Insert[int, float64](2147483647, math.Inf(0))
	test.Insert[int, float64](2147483647, math.Inf(-1))
	test.Insert[int, *int](0, nil)
	test.Insert[int, *int](-2147483648, nil)
	test.Insert[int, *int](2147483647, nil)
	test.Insert[int, bool](0, true)
	test.Insert[int, bool](0, false)
	test.Insert[int, bool](-2147483648, true)
	test.Insert[int, bool](-2147483648, false)
	test.Insert[int, bool](2147483647, true)
	test.Insert[int, bool](2147483647, false)
	test.Insert[int, string](0, "")
	test.Insert[int, string](-2147483648, "")
	test.Insert[int, string](2147483647, "")
	test.Insert[int, int](0, 0)
	test.Insert[int, int](0, -2147483648)
	test.Insert[int, int](0, 2147483647)
	test.Insert[int, int](-2147483648, 0)
	test.Insert[int, int](-2147483648, -2147483648)
	test.Insert[int, int](-2147483648, 2147483647)
	test.Insert[int, int](2147483647, 0)
	test.Insert[int, int](2147483647, -2147483648)
	test.Insert[int, int](2147483647, 2147483647)
	test.Insert[int, int64](0, 0)
	test.Insert[int, int64](0, -9223372036854775808)
	test.Insert[int, int64](0, 9223372036854775807)
	test.Insert[int, int64](-2147483648, 0)
	test.Insert[int, int64](-2147483648, -9223372036854775808)
	test.Insert[int, int64](-2147483648, 9223372036854775807)
	test.Insert[int, int64](2147483647, 0)
	test.Insert[int, int64](2147483647, -9223372036854775808)
	test.Insert[int, int64](2147483647, 9223372036854775807)
	test.Insert[int, []int](0, []int{})
	test.Insert[int, []int](0, nil)
	test.Insert[int, []int](-2147483648, []int{})
	test.Insert[int, []int](-2147483648, nil)
	test.Insert[int, []int](2147483647, []int{})
	test.Insert[int, []int](2147483647, nil)
	test.Insert[int, struct{}](0, struct{}{})
	test.Insert[int, struct{}](-2147483648, struct{}{})
	test.Insert[int, struct{}](2147483647, struct{}{})
	test.Insert[int64, struct{}](0, struct{}{})
	test.Insert[int64, struct{}](-9223372036854775808, struct{}{})
	test.Insert[int64, struct{}](9223372036854775807, struct{}{})
	test.Insert[int64, int](0, 0)
	test.Insert[int64, int](0, -2147483648)
	test.Insert[int64, int](0, 2147483647)
	test.Insert[int64, int](-9223372036854775808, 0)
	test.Insert[int64, int](-9223372036854775808, -2147483648)
	test.Insert[int64, int](-9223372036854775808, 2147483647)
	test.Insert[int64, int](9223372036854775807, 0)
	test.Insert[int64, int](9223372036854775807, -2147483648)
	test.Insert[int64, int](9223372036854775807, 2147483647)
	test.Insert[int64, int64](0, 0)
	test.Insert[int64, int64](0, -9223372036854775808)
	test.Insert[int64, int64](0, 9223372036854775807)
	test.Insert[int64, int64](-9223372036854775808, 0)
	test.Insert[int64, int64](-9223372036854775808, -9223372036854775808)
	test.Insert[int64, int64](-9223372036854775808, 9223372036854775807)
	test.Insert[int64, int64](9223372036854775807, 0)
	test.Insert[int64, int64](9223372036854775807, -9223372036854775808)
	test.Insert[int64, int64](9223372036854775807, 9223372036854775807)
	test.Insert[int64, uint](0, 0)
	test.Insert[int64, uint](0, 18446744073709551615)
	test.Insert[int64, uint](-9223372036854775808, 0)
	test.Insert[int64, uint](-9223372036854775808, 18446744073709551615)
	test.Insert[int64, uint](9223372036854775807, 0)
	test.Insert[int64, uint](9223372036854775807, 18446744073709551615)
	test.Insert[int64, float32](0, 0)
	test.Insert[int64, float32](0, float32(math.NaN()))
	test.Insert[int64, float32](0, float32(math.Inf(0)))
	test.Insert[int64, float32](0, float32(math.Inf(-1)))
	test.Insert[int64, float32](-9223372036854775808, 0)
	test.Insert[int64, float32](-9223372036854775808, float32(math.NaN()))
	test.Insert[int64, float32](-9223372036854775808, float32(math.Inf(0)))
	test.Insert[int64, float32](-9223372036854775808, float32(math.Inf(-1)))
	test.Insert[int64, float32](9223372036854775807, 0)
	test.Insert[int64, float32](9223372036854775807, float32(math.NaN()))
	test.Insert[int64, float32](9223372036854775807, float32(math.Inf(0)))
	test.Insert[int64, float32](9223372036854775807, float32(math.Inf(-1)))
	test.Insert[int64, float64](0, 0)
	test.Insert[int64, float64](0, math.NaN())
	test.Insert[int64, float64](0, math.Inf(0))
	test.Insert[int64, float64](0, math.Inf(-1))
	test.Insert[int64, float64](-9223372036854775808, 0)
	test.Insert[int64, float64](-9223372036854775808, math.NaN())
	test.Insert[int64, float64](-9223372036854775808, math.Inf(0))
	test.Insert[int64, float64](-9223372036854775808, math.Inf(-1))
	test.Insert[int64, float64](9223372036854775807, 0)
	test.Insert[int64, float64](9223372036854775807, math.NaN())
	test.Insert[int64, float64](9223372036854775807, math.Inf(0))
	test.Insert[int64, float64](9223372036854775807, math.Inf(-1))
	test.Insert[int64, *int](0, nil)
	test.Insert[int64, *int](-9223372036854775808, nil)
	test.Insert[int64, *int](9223372036854775807, nil)
	test.Insert[int64, bool](0, true)
	test.Insert[int64, bool](0, false)
	test.Insert[int64, bool](-9223372036854775808, true)
	test.Insert[int64, bool](-9223372036854775808, false)
	test.Insert[int64, bool](9223372036854775807, true)
	test.Insert[int64, bool](9223372036854775807, false)
	test.Insert[int64, string](0, "")
	test.Insert[int64, string](-9223372036854775808, "")
	test.Insert[int64, string](9223372036854775807, "")
	test.Insert[int64, []int](0, []int{})
	test.Insert[int64, []int](0, nil)
	test.Insert[int64, []int](-9223372036854775808, []int{})
	test.Insert[int64, []int](-9223372036854775808, nil)
	test.Insert[int64, []int](9223372036854775807, []int{})
	test.Insert[int64, []int](9223372036854775807, nil)
	test.Insert[uint, int](0, 0)
	test.Insert[uint, int](0, -2147483648)
	test.Insert[uint, int](0, 2147483647)
	test.Insert[uint, int](18446744073709551615, 0)
	test.Insert[uint, int](18446744073709551615, -2147483648)
	test.Insert[uint, int](18446744073709551615, 2147483647)
	test.Insert[uint, int64](0, 0)
	test.Insert[uint, int64](0, -9223372036854775808)
	test.Insert[uint, int64](0, 9223372036854775807)
	test.Insert[uint, int64](18446744073709551615, 0)
	test.Insert[uint, int64](18446744073709551615, -9223372036854775808)
	test.Insert[uint, int64](18446744073709551615, 9223372036854775807)
	test.Insert[uint, uint](0, 0)
	test.Insert[uint, uint](0, 18446744073709551615)
	test.Insert[uint, uint](18446744073709551615, 0)
	test.Insert[uint, uint](18446744073709551615, 18446744073709551615)
	test.Insert[uint, float32](0, 0)
	test.Insert[uint, float32](0, float32(math.NaN()))
	test.Insert[uint, float32](0, float32(math.Inf(0)))
	test.Insert[uint, float32](0, float32(math.Inf(-1)))
	test.Insert[uint, float32](18446744073709551615, 0)
	test.Insert[uint, float32](18446744073709551615, float32(math.NaN()))
	test.Insert[uint, float32](18446744073709551615, float32(math.Inf(0)))
	test.Insert[uint, float32](18446744073709551615, float32(math.Inf(-1)))
	test.Insert[uint, float64](0, 0)
	test.Insert[uint, float64](0, math.NaN())
	test.Insert[uint, float64](0, math.Inf(0))
	test.Insert[uint, float64](0, math.Inf(-1))
	test.Insert[uint, float64](18446744073709551615, 0)
	test.Insert[uint, float64](18446744073709551615, math.NaN())
	test.Insert[uint, float64](18446744073709551615, math.Inf(0))
	test.Insert[uint, float64](18446744073709551615, math.Inf(-1))
	test.Insert[uint, *int](0, nil)
	test.Insert[uint, *int](18446744073709551615, nil)
	test.Insert[uint, bool](0, true)
	test.Insert[uint, bool](0, false)
	test.Insert[uint, bool](18446744073709551615, true)
	test.Insert[uint, bool](18446744073709551615, false)
	test.Insert[uint, string](0, "")
	test.Insert[uint, string](18446744073709551615, "")
	test.Insert[uint, []int](0, []int{})
	test.Insert[uint, []int](0, nil)
	test.Insert[uint, []int](18446744073709551615, []int{})
	test.Insert[uint, []int](18446744073709551615, nil)
	test.Insert[uint, struct{}](0, struct{}{})
	test.Insert[uint, struct{}](18446744073709551615, struct{}{})
	test.Insert[float32, struct{}](0, struct{}{})
	test.Insert[float32, struct{}](float32(math.NaN()), struct{}{})
	test.Insert[float32, struct{}](float32(math.Inf(0)), struct{}{})
	test.Insert[float32, struct{}](float32(math.Inf(-1)), struct{}{})
	test.Insert[float32, string](0, "")
	test.Insert[float32, string](float32(math.NaN()), "")
	test.Insert[float32, string](float32(math.Inf(0)), "")
	test.Insert[float32, string](float32(math.Inf(-1)), "")
	test.Insert[float32, int](0, 0)
	test.Insert[float32, int](0, -2147483648)
	test.Insert[float32, int](0, 2147483647)
	test.Insert[float32, int](float32(math.NaN()), 0)
	test.Insert[float32, int](float32(math.NaN()), -2147483648)
	test.Insert[float32, int](float32(math.NaN()), 2147483647)
	test.Insert[float32, int](float32(math.Inf(0)), 0)
	test.Insert[float32, int](float32(math.Inf(0)), -2147483648)
	test.Insert[float32, int](float32(math.Inf(0)), 2147483647)
	test.Insert[float32, int](float32(math.Inf(-1)), 0)
	test.Insert[float32, int](float32(math.Inf(-1)), -2147483648)
	test.Insert[float32, int](float32(math.Inf(-1)), 2147483647)
	test.Insert[float32, int64](0, 0)
	test.Insert[float32, int64](0, -9223372036854775808)
	test.Insert[float32, int64](0, 9223372036854775807)
	test.Insert[float32, int64](float32(math.NaN()), 0)
	test.Insert[float32, int64](float32(math.NaN()), -9223372036854775808)
	test.Insert[float32, int64](float32(math.NaN()), 9223372036854775807)
	test.Insert[float32, int64](float32(math.Inf(0)), 0)
	test.Insert[float32, int64](float32(math.Inf(0)), -9223372036854775808)
	test.Insert[float32, int64](float32(math.Inf(0)), 9223372036854775807)
	test.Insert[float32, int64](float32(math.Inf(-1)), 0)
	test.Insert[float32, int64](float32(math.Inf(-1)), -9223372036854775808)
	test.Insert[float32, int64](float32(math.Inf(-1)), 9223372036854775807)
	test.Insert[float32, uint](0, 0)
	test.Insert[float32, uint](0, 18446744073709551615)
	test.Insert[float32, uint](float32(math.NaN()), 0)
	test.Insert[float32, uint](float32(math.NaN()), 18446744073709551615)
	test.Insert[float32, uint](float32(math.Inf(0)), 0)
	test.Insert[float32, uint](float32(math.Inf(0)), 18446744073709551615)
	test.Insert[float32, uint](float32(math.Inf(-1)), 0)
	test.Insert[float32, uint](float32(math.Inf(-1)), 18446744073709551615)
	test.Insert[float32, float32](0, 0)
	test.Insert[float32, float32](0, float32(math.NaN()))
	test.Insert[float32, float32](0, float32(math.Inf(0)))
	test.Insert[float32, float32](0, float32(math.Inf(-1)))
	test.Insert[float32, float32](float32(math.NaN()), 0)
	test.Insert[float32, float32](float32(math.NaN()), float32(math.NaN()))
	test.Insert[float32, float32](float32(math.NaN()), float32(math.Inf(0)))
	test.Insert[float32, float32](float32(math.NaN()), float32(math.Inf(-1)))
	test.Insert[float32, float32](float32(math.Inf(0)), 0)
	test.Insert[float32, float32](float32(math.Inf(0)), float32(math.NaN()))
	test.Insert[float32, float32](float32(math.Inf(0)), float32(math.Inf(0)))
	test.Insert[float32, float32](float32(math.Inf(0)), float32(math.Inf(-1)))
	test.Insert[float32, float32](float32(math.Inf(-1)), 0)
	test.Insert[float32, float32](float32(math.Inf(-1)), float32(math.NaN()))
	test.Insert[float32, float32](float32(math.Inf(-1)), float32(math.Inf(0)))
	test.Insert[float32, float32](float32(math.Inf(-1)), float32(math.Inf(-1)))
	test.Insert[float32, float64](0, 0)
	test.Insert[float32, float64](0, math.NaN())
	test.Insert[float32, float64](0, math.Inf(0))
	test.Insert[float32, float64](0, math.Inf(-1))
	test.Insert[float32, float64](float32(math.NaN()), 0)
	test.Insert[float32, float64](float32(math.NaN()), math.NaN())
	test.Insert[float32, float64](float32(math.NaN()), math.Inf(0))
	test.Insert[float32, float64](float32(math.NaN()), math.Inf(-1))
	test.Insert[float32, float64](float32(math.Inf(0)), 0)
	test.Insert[float32, float64](float32(math.Inf(0)), math.NaN())
	test.Insert[float32, float64](float32(math.Inf(0)), math.Inf(0))
	test.Insert[float32, float64](float32(math.Inf(0)), math.Inf(-1))
	test.Insert[float32, float64](float32(math.Inf(-1)), 0)
	test.Insert[float32, float64](float32(math.Inf(-1)), math.NaN())
	test.Insert[float32, float64](float32(math.Inf(-1)), math.Inf(0))
	test.Insert[float32, float64](float32(math.Inf(-1)), math.Inf(-1))
	test.Insert[float32, *int](0, nil)
	test.Insert[float32, *int](float32(math.NaN()), nil)
	test.Insert[float32, *int](float32(math.Inf(0)), nil)
	test.Insert[float32, *int](float32(math.Inf(-1)), nil)
	test.Insert[float32, bool](0, true)
	test.Insert[float32, bool](0, false)
	test.Insert[float32, bool](float32(math.NaN()), true)
	test.Insert[float32, bool](float32(math.NaN()), false)
	test.Insert[float32, bool](float32(math.Inf(0)), true)
	test.Insert[float32, bool](float32(math.Inf(0)), false)
	test.Insert[float32, bool](float32(math.Inf(-1)), true)
	test.Insert[float32, bool](float32(math.Inf(-1)), false)
	test.Insert[float32, []int](0, []int{})
	test.Insert[float32, []int](0, nil)
	test.Insert[float32, []int](float32(math.NaN()), []int{})
	test.Insert[float32, []int](float32(math.NaN()), nil)
	test.Insert[float32, []int](float32(math.Inf(0)), []int{})
	test.Insert[float32, []int](float32(math.Inf(0)), nil)
	test.Insert[float32, []int](float32(math.Inf(-1)), []int{})
	test.Insert[float32, []int](float32(math.Inf(-1)), nil)
	test.Insert[float64, struct{}](0, struct{}{})
	test.Insert[float64, struct{}](math.NaN(), struct{}{})
	test.Insert[float64, struct{}](math.Inf(0), struct{}{})
	test.Insert[float64, struct{}](math.Inf(-1), struct{}{})
	test.Insert[float64, int64](0, 0)
	test.Insert[float64, int64](0, -9223372036854775808)
	test.Insert[float64, int64](0, 9223372036854775807)
	test.Insert[float64, int64](math.NaN(), 0)
	test.Insert[float64, int64](math.NaN(), -9223372036854775808)
	test.Insert[float64, int64](math.NaN(), 9223372036854775807)
	test.Insert[float64, int64](math.Inf(0), 0)
	test.Insert[float64, int64](math.Inf(0), -9223372036854775808)
	test.Insert[float64, int64](math.Inf(0), 9223372036854775807)
	test.Insert[float64, int64](math.Inf(-1), 0)
	test.Insert[float64, int64](math.Inf(-1), -9223372036854775808)
	test.Insert[float64, int64](math.Inf(-1), 9223372036854775807)
	test.Insert[float64, uint](0, 0)
	test.Insert[float64, uint](0, 18446744073709551615)
	test.Insert[float64, uint](math.NaN(), 0)
	test.Insert[float64, uint](math.NaN(), 18446744073709551615)
	test.Insert[float64, uint](math.Inf(0), 0)
	test.Insert[float64, uint](math.Inf(0), 18446744073709551615)
	test.Insert[float64, uint](math.Inf(-1), 0)
	test.Insert[float64, uint](math.Inf(-1), 18446744073709551615)
	test.Insert[float64, float32](0, 0)
	test.Insert[float64, float32](0, float32(math.NaN()))
	test.Insert[float64, float32](0, float32(math.Inf(0)))
	test.Insert[float64, float32](0, float32(math.Inf(-1)))
	test.Insert[float64, float32](math.NaN(), 0)
	test.Insert[float64, float32](math.NaN(), float32(math.NaN()))
	test.Insert[float64, float32](math.NaN(), float32(math.Inf(0)))
	test.Insert[float64, float32](math.NaN(), float32(math.Inf(-1)))
	test.Insert[float64, float32](math.Inf(0), 0)
	test.Insert[float64, float32](math.Inf(0), float32(math.NaN()))
	test.Insert[float64, float32](math.Inf(0), float32(math.Inf(0)))
	test.Insert[float64, float32](math.Inf(0), float32(math.Inf(-1)))
	test.Insert[float64, float32](math.Inf(-1), 0)
	test.Insert[float64, float32](math.Inf(-1), float32(math.NaN()))
	test.Insert[float64, float32](math.Inf(-1), float32(math.Inf(0)))
	test.Insert[float64, float32](math.Inf(-1), float32(math.Inf(-1)))
	test.Insert[float64, float64](0, 0)
	test.Insert[float64, float64](0, math.NaN())
	test.Insert[float64, float64](0, math.Inf(0))
	test.Insert[float64, float64](0, math.Inf(-1))
	test.Insert[float64, float64](math.NaN(), 0)
	test.Insert[float64, float64](math.NaN(), math.NaN())
	test.Insert[float64, float64](math.NaN(), math.Inf(0))
	test.Insert[float64, float64](math.NaN(), math.Inf(-1))
	test.Insert[float64, float64](math.Inf(0), 0)
	test.Insert[float64, float64](math.Inf(0), math.NaN())
	test.Insert[float64, float64](math.Inf(0), math.Inf(0))
	test.Insert[float64, float64](math.Inf(0), math.Inf(-1))
	test.Insert[float64, float64](math.Inf(-1), 0)
	test.Insert[float64, float64](math.Inf(-1), math.NaN())
	test.Insert[float64, float64](math.Inf(-1), math.Inf(0))
	test.Insert[float64, float64](math.Inf(-1), math.Inf(-1))
	test.Insert[float64, *int](0, nil)
	test.Insert[float64, *int](math.NaN(), nil)
	test.Insert[float64, *int](math.Inf(0), nil)
	test.Insert[float64, *int](math.Inf(-1), nil)
	test.Insert[float64, bool](0, true)
	test.Insert[float64, bool](0, false)
	test.Insert[float64, bool](math.NaN(), true)
	test.Insert[float64, bool](math.NaN(), false)
	test.Insert[float64, bool](math.Inf(0), true)
	test.Insert[float64, bool](math.Inf(0), false)
	test.Insert[float64, bool](math.Inf(-1), true)
	test.Insert[float64, bool](math.Inf(-1), false)
	test.Insert[float64, string](0, "")
	test.Insert[float64, string](math.NaN(), "")
	test.Insert[float64, string](math.Inf(0), "")
	test.Insert[float64, string](math.Inf(-1), "")
	test.Insert[float64, int](0, 0)
	test.Insert[float64, int](0, -2147483648)
	test.Insert[float64, int](0, 2147483647)
	test.Insert[float64, int](math.NaN(), 0)
	test.Insert[float64, int](math.NaN(), -2147483648)
	test.Insert[float64, int](math.NaN(), 2147483647)
	test.Insert[float64, int](math.Inf(0), 0)
	test.Insert[float64, int](math.Inf(0), -2147483648)
	test.Insert[float64, int](math.Inf(0), 2147483647)
	test.Insert[float64, int](math.Inf(-1), 0)
	test.Insert[float64, int](math.Inf(-1), -2147483648)
	test.Insert[float64, int](math.Inf(-1), 2147483647)
	test.Insert[float64, []int](0, []int{})
	test.Insert[float64, []int](0, nil)
	test.Insert[float64, []int](math.NaN(), []int{})
	test.Insert[float64, []int](math.NaN(), nil)
	test.Insert[float64, []int](math.Inf(0), []int{})
	test.Insert[float64, []int](math.Inf(0), nil)
	test.Insert[float64, []int](math.Inf(-1), []int{})
	test.Insert[float64, []int](math.Inf(-1), nil)
	test.Insert[*int, string](nil, "")
	test.Insert[*int, int](nil, 0)
	test.Insert[*int, int](nil, -2147483648)
	test.Insert[*int, int](nil, 2147483647)
	test.Insert[*int, int64](nil, 0)
	test.Insert[*int, int64](nil, -9223372036854775808)
	test.Insert[*int, int64](nil, 9223372036854775807)
	test.Insert[*int, uint](nil, 0)
	test.Insert[*int, uint](nil, 18446744073709551615)
	test.Insert[*int, float32](nil, 0)
	test.Insert[*int, float32](nil, float32(math.NaN()))
	test.Insert[*int, float32](nil, float32(math.Inf(0)))
	test.Insert[*int, float32](nil, float32(math.Inf(-1)))
	test.Insert[*int, float64](nil, 0)
	test.Insert[*int, float64](nil, math.NaN())
	test.Insert[*int, float64](nil, math.Inf(0))
	test.Insert[*int, float64](nil, math.Inf(-1))
	test.Insert[*int, *int](nil, nil)
	test.Insert[*int, bool](nil, true)
	test.Insert[*int, bool](nil, false)
	test.Insert[*int, []int](nil, []int{})
	test.Insert[*int, []int](nil, nil)
	test.Insert[*int, struct{}](nil, struct{}{})
	test.Insert[bool, int64](true, 0)
	test.Insert[bool, int64](true, -9223372036854775808)
	test.Insert[bool, int64](true, 9223372036854775807)
	test.Insert[bool, int64](false, 0)
	test.Insert[bool, int64](false, -9223372036854775808)
	test.Insert[bool, int64](false, 9223372036854775807)
	test.Insert[bool, uint](true, 0)
	test.Insert[bool, uint](true, 18446744073709551615)
	test.Insert[bool, uint](false, 0)
	test.Insert[bool, uint](false, 18446744073709551615)
	test.Insert[bool, float32](true, 0)
	test.Insert[bool, float32](true, float32(math.NaN()))
	test.Insert[bool, float32](true, float32(math.Inf(0)))
	test.Insert[bool, float32](true, float32(math.Inf(-1)))
	test.Insert[bool, float32](false, 0)
	test.Insert[bool, float32](false, float32(math.NaN()))
	test.Insert[bool, float32](false, float32(math.Inf(0)))
	test.Insert[bool, float32](false, float32(math.Inf(-1)))
	test.Insert[bool, float64](true, 0)
	test.Insert[bool, float64](true, math.NaN())
	test.Insert[bool, float64](true, math.Inf(0))
	test.Insert[bool, float64](true, math.Inf(-1))
	test.Insert[bool, float64](false, 0)
	test.Insert[bool, float64](false, math.NaN())
	test.Insert[bool, float64](false, math.Inf(0))
	test.Insert[bool, float64](false, math.Inf(-1))
	test.Insert[bool, *int](true, nil)
	test.Insert[bool, *int](false, nil)
	test.Insert[bool, bool](true, true)
	test.Insert[bool, bool](true, false)
	test.Insert[bool, bool](false, true)
	test.Insert[bool, bool](false, false)
	test.Insert[bool, string](true, "")
	test.Insert[bool, string](false, "")
	test.Insert[bool, int](true, 0)
	test.Insert[bool, int](true, -2147483648)
	test.Insert[bool, int](true, 2147483647)
	test.Insert[bool, int](false, 0)
	test.Insert[bool, int](false, -2147483648)
	test.Insert[bool, int](false, 2147483647)
	test.Insert[bool, []int](true, []int{})
	test.Insert[bool, []int](true, nil)
	test.Insert[bool, []int](false, []int{})
	test.Insert[bool, []int](false, nil)
	test.Insert[bool, struct{}](true, struct{}{})
	test.Insert[bool, struct{}](false, struct{}{})
	test.Insert[[]int, struct{}]([]int{}, struct{}{})
	test.Insert[[]int, struct{}](nil, struct{}{})
	test.Insert[[]int, float64]([]int{}, 0)
	test.Insert[[]int, float64]([]int{}, math.NaN())
	test.Insert[[]int, float64]([]int{}, math.Inf(0))
	test.Insert[[]int, float64]([]int{}, math.Inf(-1))
	test.Insert[[]int, float64](nil, 0)
	test.Insert[[]int, float64](nil, math.NaN())
	test.Insert[[]int, float64](nil, math.Inf(0))
	test.Insert[[]int, float64](nil, math.Inf(-1))
	test.Insert[[]int, *int]([]int{}, nil)
	test.Insert[[]int, *int](nil, nil)
	test.Insert[[]int, bool]([]int{}, true)
	test.Insert[[]int, bool]([]int{}, false)
	test.Insert[[]int, bool](nil, true)
	test.Insert[[]int, bool](nil, false)
	test.Insert[[]int, string]([]int{}, "")
	test.Insert[[]int, string](nil, "")
	test.Insert[[]int, int]([]int{}, 0)
	test.Insert[[]int, int]([]int{}, -2147483648)
	test.Insert[[]int, int]([]int{}, 2147483647)
	test.Insert[[]int, int](nil, 0)
	test.Insert[[]int, int](nil, -2147483648)
	test.Insert[[]int, int](nil, 2147483647)
	test.Insert[[]int, int64]([]int{}, 0)
	test.Insert[[]int, int64]([]int{}, -9223372036854775808)
	test.Insert[[]int, int64]([]int{}, 9223372036854775807)
	test.Insert[[]int, int64](nil, 0)
	test.Insert[[]int, int64](nil, -9223372036854775808)
	test.Insert[[]int, int64](nil, 9223372036854775807)
	test.Insert[[]int, uint]([]int{}, 0)
	test.Insert[[]int, uint]([]int{}, 18446744073709551615)
	test.Insert[[]int, uint](nil, 0)
	test.Insert[[]int, uint](nil, 18446744073709551615)
	test.Insert[[]int, float32]([]int{}, 0)
	test.Insert[[]int, float32]([]int{}, float32(math.NaN()))
	test.Insert[[]int, float32]([]int{}, float32(math.Inf(0)))
	test.Insert[[]int, float32]([]int{}, float32(math.Inf(-1)))
	test.Insert[[]int, float32](nil, 0)
	test.Insert[[]int, float32](nil, float32(math.NaN()))
	test.Insert[[]int, float32](nil, float32(math.Inf(0)))
	test.Insert[[]int, float32](nil, float32(math.Inf(-1)))
	test.Insert[[]int, []int]([]int{}, []int{})
	test.Insert[[]int, []int]([]int{}, nil)
	test.Insert[[]int, []int](nil, []int{})
	test.Insert[[]int, []int](nil, nil)
	test.Insert[struct{}, bool](struct{}{}, true)
	test.Insert[struct{}, bool](struct{}{}, false)
	test.Insert[struct{}, string](struct{}{}, "")
	test.Insert[struct{}, int](struct{}{}, 0)
	test.Insert[struct{}, int](struct{}{}, -2147483648)
	test.Insert[struct{}, int](struct{}{}, 2147483647)
	test.Insert[struct{}, int64](struct{}{}, 0)
	test.Insert[struct{}, int64](struct{}{}, -9223372036854775808)
	test.Insert[struct{}, int64](struct{}{}, 9223372036854775807)
	test.Insert[struct{}, uint](struct{}{}, 0)
	test.Insert[struct{}, uint](struct{}{}, 18446744073709551615)
	test.Insert[struct{}, float32](struct{}{}, 0)
	test.Insert[struct{}, float32](struct{}{}, float32(math.NaN()))
	test.Insert[struct{}, float32](struct{}{}, float32(math.Inf(0)))
	test.Insert[struct{}, float32](struct{}{}, float32(math.Inf(-1)))
	test.Insert[struct{}, float64](struct{}{}, 0)
	test.Insert[struct{}, float64](struct{}{}, math.NaN())
	test.Insert[struct{}, float64](struct{}{}, math.Inf(0))
	test.Insert[struct{}, float64](struct{}{}, math.Inf(-1))
	test.Insert[struct{}, *int](struct{}{}, nil)
	test.Insert[struct{}, []int](struct{}{}, []int{})
	test.Insert[struct{}, []int](struct{}{}, nil)
	test.Insert[struct{}, struct{}](struct{}{}, struct{}{})

}
