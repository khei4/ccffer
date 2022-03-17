package test

type X interface {
	~int | ~string
}

// func F[T X](v T) {
// 	println(v + v)
// 	println(string(v))
// 	println(any(v).(int))
// }
// func main() {
// 	F(42)
// 	F("42")
// }

func F(v int, s string) string {
	return s
}

func main() {
	F(32, "hoge")
}
