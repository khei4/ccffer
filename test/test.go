package test

import "strconv"

type X interface {
	~int | ~string
}

func FG[T X, Y any](v T) {
	println(v + v)
	println(string(v))
	println(any(v).(int))
}

type S struct {
	hoge int
	fuga string
}

func F(sli []int, stct S, p *int) string {
	if sli == nil {
		return "hoge"
	}
	x := sli[0]
	return strconv.Itoa(x)
}

func main() {
	x := 12
	F([]int{32}, S{hoge: 12, fuga: "hoge"}, &x)
	FG[int, string](42)
	FG[string, string]("42")
}
