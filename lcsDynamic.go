package main

import (
	"fmt"
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func fillDynMatrix(x, y []byte) [][]int {
	L := make([][]int, len(x))
	for i:=0; i < len(x); i++ {
		L[i] = make([]int, len(y))
	}

	for i:= 1; i < len(x); i++ {
		for j:=1; j < len(y); j++ {
			if x[i] == y[j]{
				L[i][j] = L[i -1][j -1] + 1
			} else {
				L[i][j] = Max(Max(L[i - 1][j], L[i][j - 1]), L[i -1][j -1])
			}
		}
	}
	return L
}

func lcsDyn(L [][]int, x , y []byte) string {
	res := ""
	k := -1
	for i:= 0; i < len(L); i++ {
		for j:= 0; j < len(L[0]);j++ {
			if L[i][j] > k {
				res += string(x[i])
				k += 1
			}
		}
	}
	return res
}

func main() {
	x := []byte("abc")
	y := []byte("abcdef")
	
	L := fillDynMatrix(x, y)
	fmt.Println(lcsDyn(L, x, y))
}
