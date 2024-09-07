package main

func CollatzStep(number int) int {
	if number%2 == 0 {
		return number / 2
	}
	return number*3 + 1
}
