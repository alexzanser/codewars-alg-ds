package kata

func Well(x []string) string {
	var num int

	for i := 0; i < len(x); i++ {
		if x[i] == "good" {
			num += 1
		}
	}

	if num == 2 || num == 1 {
		return "Publish!"
	} else if num >= 2 {
		return "I smell a series!"
	} else {
		return "Fail!"
	}
}
