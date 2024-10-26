package main

func TwoSums(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j {
		_sum := numbers[i] + numbers[j]

		if _sum == target {
			return []int{i + 1, j + 1}
		}

		if _sum > target {
			j = j - 1
		} else {
			i = i + 1
		}
	}
	return []int{}
}
