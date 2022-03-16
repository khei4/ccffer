package main

import (
	"fmt"
)

func SignleArg(a int) {
	fmt.Print(a)
}
func SignleRes() int {
	return 42
}
func NamedRes() (g int) {
	g = 42
	return
}

func main() {
	SignleArg(13)
	_ = SignleRes()
	_ = NamedRes()
}
