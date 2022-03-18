package test_test

import (
	"ccffer/test"
	"math"
	"testing"
)

func TestFG(t *testing.T) {
	test.FG[string, uint]("")
	test.FG[string, uint]("hogehogehogehoge")
	test.FG[string, []int]("")
	test.FG[string, []int]("hogehogehogehoge")
	test.FG[string, *int]("")
	test.FG[string, *int]("hogehogehogehoge")
	test.FG[string, bool]("")
	test.FG[string, bool]("hogehogehogehoge")
	test.FG[string, string]("")
	test.FG[string, string]("hogehogehogehoge")
	test.FG[string, int]("")
	test.FG[string, int]("hogehogehogehoge")
	test.FG[string, int64]("")
	test.FG[string, int64]("hogehogehogehoge")
	test.FG[string, struct{}]("")
	test.FG[string, struct{}]("hogehogehogehoge")
	test.FG[string, float32]("")
	test.FG[string, float32]("hogehogehogehoge")
	test.FG[string, float64]("")
	test.FG[string, float64]("hogehogehogehoge")
	test.FG[int, float32](0)
	test.FG[int, float32](-2147483648)
	test.FG[int, float32](2147483647)
	test.FG[int, float64](0)
	test.FG[int, float64](-2147483648)
	test.FG[int, float64](2147483647)
	test.FG[int, *int](0)
	test.FG[int, *int](-2147483648)
	test.FG[int, *int](2147483647)
	test.FG[int, bool](0)
	test.FG[int, bool](-2147483648)
	test.FG[int, bool](2147483647)
	test.FG[int, string](0)
	test.FG[int, string](-2147483648)
	test.FG[int, string](2147483647)
	test.FG[int, int](0)
	test.FG[int, int](-2147483648)
	test.FG[int, int](2147483647)
	test.FG[int, int64](0)
	test.FG[int, int64](-2147483648)
	test.FG[int, int64](2147483647)
	test.FG[int, struct{}](0)
	test.FG[int, struct{}](-2147483648)
	test.FG[int, struct{}](2147483647)
	test.FG[int, uint](0)
	test.FG[int, uint](-2147483648)
	test.FG[int, uint](2147483647)
	test.FG[int, []int](0)
	test.FG[int, []int](-2147483648)
	test.FG[int, []int](2147483647)

}
func TestF(t *testing.T) {
	test.F(0, "", 0)
	test.F(0, "", math.NaN())
	test.F(0, "", math.Inf(0))
	test.F(0, "", math.Inf(-1))
	test.F(0, "hogehogehogehoge", 0)
	test.F(0, "hogehogehogehoge", math.NaN())
	test.F(0, "hogehogehogehoge", math.Inf(0))
	test.F(0, "hogehogehogehoge", math.Inf(-1))
	test.F(-2147483648, "", 0)
	test.F(-2147483648, "", math.NaN())
	test.F(-2147483648, "", math.Inf(0))
	test.F(-2147483648, "", math.Inf(-1))
	test.F(-2147483648, "hogehogehogehoge", 0)
	test.F(-2147483648, "hogehogehogehoge", math.NaN())
	test.F(-2147483648, "hogehogehogehoge", math.Inf(0))
	test.F(-2147483648, "hogehogehogehoge", math.Inf(-1))
	test.F(2147483647, "", 0)
	test.F(2147483647, "", math.NaN())
	test.F(2147483647, "", math.Inf(0))
	test.F(2147483647, "", math.Inf(-1))
	test.F(2147483647, "hogehogehogehoge", 0)
	test.F(2147483647, "hogehogehogehoge", math.NaN())
	test.F(2147483647, "hogehogehogehoge", math.Inf(0))
	test.F(2147483647, "hogehogehogehoge", math.Inf(-1))

}
func Testmain(t *testing.T) {

}
