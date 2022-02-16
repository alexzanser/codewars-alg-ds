package kata

func MaximumSubarraySum(numbers []int) int {
	subsum := 0
	finalMax := 0

	for _, val := range numbers {
		if subsum >= 0 {
			subsum += val
		} else {
			subsum = val
		}
		if subsum > finalMax {
			finalMax = subsum
		}
	}
	return finalMax
}
