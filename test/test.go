package test

var m map[interface{}]interface{}

func Insert[K, V any](k K, v V) {
	m[k] = v
}

// func Read(s []int) int {
// 	a := s[0]
// 	return a
// }

// type Point struct {
// 	x int
// 	y int
// }
// func Addxy(p *Point) int {
// 	return p.x + p.y
// }

// type S struct {
// 	hoge int
// 	fuga string
// }

// func F(sli []int, stct S, p *int) string {
// 	if len(sli) == 0 {
// 		return "hoge"
// 	}
// 	x := sli[0]
// 	return strconv.Itoa(x)
// }

func main() {
}
