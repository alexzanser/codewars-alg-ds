package kata

func Parse(data string) []int {
	val := 0
	arr := []int{}

	for _, c := range data {
		switch string(c) {
		case "i":
			val += 1
		case "d":
			val -= 1
		case "s":
			val = val * val
		case "o":
			arr = append(arr, val)
		}
	}
	return arr
}
