package kata

import "sort"

func IsTriangle(a, b, c int) bool {
	d := []int{a, b, c}
	sort.Ints(d)
	max := d[2]
	if max < a+b+c-max {
		return true
	}
	return false
}
