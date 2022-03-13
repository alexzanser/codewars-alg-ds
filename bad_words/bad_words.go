package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
func main() {
	var fname string
	badWords  := make(map[string]struct{})

	fmt.Scan(&fname)

	file, err := os.Open(fname)
	if err != nil {
		log.Fatal("can't open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        badWords[scanner.Text()] = struct{}{}
    }

	reader := bufio.NewReader(os.Stdin)
	for {
		line , _, err := reader.ReadLine() 
		if err != nil {
			log.Fatal(err.Error())
			break
		}
		sentence := string(line)
		if sentence == "exit" {
			fmt.Println("Bye!")
			break
		}

		words := strings.Split(sentence, " ")

		for _, word := range words {
			badword := strings.ToLower(word)
			if _, ok:= badWords[badword]; ok {
				fmt.Printf("%s ", strings.Repeat("*", len(word)))
			} else {
				fmt.Printf("%s ",word)
			}
		}
		fmt.Println()
	}
}
