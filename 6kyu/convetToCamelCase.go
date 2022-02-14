package kata

import "strings"

func ToCamelCase(s string) string {
	var output string

	for i, _ := range s {
		if i > 1 && (string(s[i-1]) == "-" || string(s[i-1]) == "_") {
			output += strings.ToUpper(string(s[i]))
		} else if string(s[i]) != "-" && string(s[i]) != "_" {
			output += string(s[i])
		}
	}
	return output
}
