package main

import (
	"fmt"
	"unicode/utf8"
)

func banner() {
	message := "Hi :)"
	width := 8
	padding := (width - utf8.RuneCountInString(message)) / 2
	for index := 0; index < padding; index++ {
		fmt.Print(" ")
	}
	fmt.Println(message)
	for index := 0; index < width; index++ {
		fmt.Print("-")
	}
	fmt.Println()
}
