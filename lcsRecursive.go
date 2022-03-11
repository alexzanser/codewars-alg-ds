package main 

import "fmt"

func lcsREC(x, y []byte) []byte {
	if len(x) == 0 || len(y) == 0{
		return nil
	}
	if x[len(x) - 1] == y[len(y) - 1] {
		return append(lcsREC(x[0:len(x) - 1], y[0:len(y) - 1]), x[len(x) - 1])
	} else {
		left := lcsREC(x[0:len(x) - 1], y)
		right:= lcsREC(x, y[0:len(y) - 1])

		if len(left) > len(right) {
			return left
		}else {
			return right
		}
	}
}

func main() {
	s1 := []byte("abcdef")
	s2 := []byte("abc")

	fmt.Println(string(lcsREC(s1, s2)))
}