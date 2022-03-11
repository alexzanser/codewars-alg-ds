package main

import "fmt"

func fillDynMatrix(x, y []byte) [][]int {
	L := make([][]int, len(x))
	for i:=0; i < len(x); i++ {
		L[i] = make([]int, len(y))
	}

	for _, line := range L {
		fmt.Println(line)
	}

	return L
}

func main() {
	x := []byte("abc")
	y := []byte("abcdef")
	
	fillDynMatrix(x, y)
}