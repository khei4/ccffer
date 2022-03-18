package test

type X interface {
	~int | ~string
}

func FG[T X, Y any](v T) {
	println(v + v)
	println(string(v))
	println(any(v).(int))
}

func F(v int, s string, f float64) string {
	return s
}

func main() {
	_ = struct{}{}
	F(32, "hoge", 0)
	FG[int, string](42)
	FG[string, string]("42")
}
