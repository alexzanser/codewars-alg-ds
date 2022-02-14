package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MoveZeros(arr []int) []int {
	l := len(arr)
	for i := l - 1; i > -1; i-- {
		if arr[i] == 0 {
			for j := i + 1; j < l; j++ {
				if arr[j] != 0 {
					arr[j], arr[j-1] = arr[j-1], arr[j]
				}
			}
		}
	}
	return arr
}

func main() {
	var digits []int

	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 10; i++ {
		digits = append(digits, int(rand.Int31()%100))
	}
	digits[8] = 0
	digits[5] = 0
	digits[3] = 0
	digits[1] = 0
	fmt.Println(digits)
	fmt.Println(MoveZeros(digits))
}
