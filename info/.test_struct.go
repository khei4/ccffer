package main

type X interface {
	~int | ~string
}

func F[T X](v T) {
	println(v + v)
	println(string(v))
	println(any(v).(int))
}
func main() {
	F(42)
	F("42")
}
