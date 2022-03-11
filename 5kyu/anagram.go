package main

import "fmt"


func setOfLetters(word string) map[string]struct{} {
	symbols := make(map[string]struct{})
	for _, l:= range []byte(word) {
		if _, ok := symbols[string(l)]; !ok {
			symbols[string(l)] = struct{}{}
		}
	}

	return symbols
}

func Anagrams(word string, words []string) []string {
	anagrams := make([]string, 0)

	symbols := setOfLetters(word)
	for _, w := range words {
		f := true
		for _, l := range w {
			if _, ok := symbols[string(l)]; !ok {
				f = false
			} 
		}
		if f {
			anagrams = append(anagrams, w)
		}
	}
	if len(anagrams) == 0{
		return nil
	}
	return anagrams
}

func main() {
	fmt.Println(Anagrams("r", []string{"crazer", "carer", "racar", "caers", "racer"}))
}