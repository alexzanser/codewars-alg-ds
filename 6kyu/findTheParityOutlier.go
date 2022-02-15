package kata

func FindOutlier(integers []int) int {
	even := []int{}
	odd := []int{}

	for _, val := range integers {
		if val%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}
	if len(even) == 1 {
		return (even[0])
	} else {
		return odd[0]
	}
}
