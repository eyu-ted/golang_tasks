package main

import (
	"fmt"
	"strings"
	"unicode"
)

func calculateFrequency(text string) {

	text = strings.ToLower(text)
	words := strings.Fields(text)
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			if !unicode.IsLetter(rune(word[j])) {
				words[i] = word[:j] + word[j+1:]
			}
		}
	}

	frequency := map[string]int{}
	for _, word := range words {
		frequency[word]++
	}
	fmt.Println(frequency)
}

func main() {
	calculateFrequency("hello's tuweuthew, sfr dF e fe?")
}
