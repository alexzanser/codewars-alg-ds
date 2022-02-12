package kata

func Points(games []string) int {
	var totalScore int

	totalScore = 0
	for i := 0; i < len(games); i++ {
		if games[i][0:1] > games[i][2:3] {
			totalScore += 3
		} else if games[i][0:1] == games[i][2:3] {
			totalScore += 1
		}
	}
	return totalScore
}
