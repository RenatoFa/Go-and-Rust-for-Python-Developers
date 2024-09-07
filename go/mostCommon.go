package main

import (
	"fmt"
	"strings"
)

func MostCommon() {
	poem := `
	those who do not feel this love pulling
        them like a river those who do not drink dawn
        like a cup of sring water or take in sunset like
        supper those who do not want to change let them sleep
	`

	frequency := make(map[string]int)

	for _, word := range strings.Fields(poem) {
		frequency[word]++
	}

	maxWord, maxNumber := "", 0

	for word, number := range frequency {
		if number > maxNumber {
			maxWord, maxNumber = word, number
		}
	}

	fmt.Println(maxWord)
}
