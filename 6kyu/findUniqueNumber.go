package kata

import "sort"

func FindUniq(arr []float32) float32 {
	arr2 := make([]float64, len(arr))
	for i, val := range arr {
		arr2[i] = float64(val)
	}
	sort.Float64s(arr2)
	if arr2[0] != arr2[1] {
		return float32(arr2[0])
	} else {
		return float32(arr2[len(arr2)-1])
	}
}
