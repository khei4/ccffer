package test

type X interface {
	~int | ~string
}

func FG[T X](v T) {
	println(v + v)
	println(string(v))
	println(any(v).(int))
}

func F(v int, s string, f float32) string {
	return s
}

func main() {
	_ = struct{}{}
	F(32, "hoge", 0)
	FG(42)
	FG("42")
}
