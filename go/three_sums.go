package main

import "slices"

func ThreeSums(numbers []int) [][]int {
	slices.Sort(numbers)
	res := [][]int{}

	for index, number := range numbers {
		if index > 0 && number == numbers[index-1] {
			continue
		}

		left, right := index+1, len(numbers)-1

		for left < right {
			three_sums := number + numbers[left] + numbers[right]

			if three_sums > 0 {
				right -= 1
			} else if three_sums < 0 {
				left += 1
			} else {
				res = append(res, []int{number, numbers[left], numbers[right]})
				left += 1
				for numbers[left] == numbers[left-1] && left < right {
					left += 1
				}
			}
		}
	}
	return res
}
