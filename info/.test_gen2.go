package main

func SignleTypeParam[T any](a T) T {
	return a
}

func TypeConstraint1[T comparable](a, b T) bool {
	return a == b
}

func main() {
	SignleTypeParam("hoge")
	SignleTypeParam(13)
	TypeConstraint1(1, 2)
	TypeConstraint1("1,2", "2,3")
}
