package main

import (
	"fmt"
	"sort"
	"unicode"
)

func CharFrequency() {
	poem := `those who do not feel this love pulling
        them like a river those who do not drink dawn
        like a cup of sring water or take in sunset like
        supper those who do not want to change let them sleep
        `

	counts := make(map[rune]int)
	count := 0.0

	for _, character := range poem {
		if unicode.IsSpace(character) {
			continue
		}
		counts[character]++
		count++
	}

	var chars []rune

	for character := range counts {
		chars = append(chars, character)
	}

	sort.Slice(chars, func(i, j int) bool {
		c1, c2 := chars[i], chars[j]
		return counts[c1] > counts[c2]
	})

	for _, c := range chars {
		n := counts[c]
		f := float64(n) / count * 100
		fmt.Printf("%c: %.2f\n", c, f)
	}
}
