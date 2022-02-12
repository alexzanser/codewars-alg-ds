package kata

import "unicode/utf8"

func RemoveChar(word string) string {
	return word[1 : utf8.RuneCountInString(word)-1]
}
