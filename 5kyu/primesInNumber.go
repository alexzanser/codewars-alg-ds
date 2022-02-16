package main

import "fmt"

func PrimeFactors(n int) string {
	dels_power := make([]int, n + 1, n + 1)
	res := ""
	for i := 2; i < len(dels_power); i++ {
		if n % i == 0 {
			for (n % i == 0 ) {
				dels_power[i] += 1
				n = n / i
			}
		}
	}
	for i := 2; i < len(dels_power); i++ {
		if dels_power[i] == 1 {
			res += fmt.Sprintf("(%d)", i)
		} else if dels_power[i] > 1 {
			res += fmt.Sprintf("(%d**%d)", i, dels_power[i])
		}
	}
	return res
}

func main() {
	x := 655580

	fmt.Println(PrimeFactors(x))
}
