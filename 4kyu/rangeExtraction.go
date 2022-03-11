package main

import (
	"fmt"
	"strconv"

)

func quickSort(digits []int) []int {
	if len(digits) < 2 {
		return digits
	}
	if len(digits) == 2 {
		if digits[0] > digits[1] {
			digits[0], digits[1] = digits[1], digits[0]
		}
		return digits
	}
	base := digits[0]
	digits = digits[1:]
	var left, right []int
	for _, elem := range digits {
		if elem <= base {
			left = append(left, elem)
		} else {
			right = append(right, elem)
		}
	}
	left = append(left, base)
	return append(quickSort(left), quickSort(right)...)
}

func nextExtra(curPos int, list []int) int {
	l := 0
	for i:=curPos + 1; i < len(list); i++ {
		if list[i] == list[i - 1] + 1 {
			l += 1
		} else {
			return l
		}
	}

	return l
}

func Solution(list []int) string {
	list = quickSort(list)
	res := ""
	for  i:= 0; i< len(list); i++ {
		next := nextExtra(i, list)
		if next < 2{
			res += strconv.Itoa(list[i]) + ","
		} else {
			res += strconv.Itoa(list[i]) + "-" + strconv.Itoa(list[i + next]) + ","
			i += next
		}
	}
	res = res[0:len(res) - 1]
	return res
}


func main() {
	list := []int{-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}
	fmt.Println(Solution(list))
} 