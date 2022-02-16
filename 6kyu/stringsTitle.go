package kata

import "strings"

func CamelCase(s string) string {
	words := strings.Split(s, " ")
	res := ""
	for _, word := range words {
		res += strings.Title(word)
	}
	return res
}
