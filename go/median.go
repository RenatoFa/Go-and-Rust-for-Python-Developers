package main

import (
	"fmt"
	"sort"
)

func media() {
	numbers := []float64{2, 1, 3}

	nums := make([]float64, len(numbers))
	copy(nums, numbers)

	sort.Float64s(nums)
	var media float64

	i := len(nums) / 2

	if len(nums)%2 == 1 {
		media = nums[i]
	} else {
		media = (nums[i-1] + nums[i]) / 2
	}
	fmt.Println(numbers)
	fmt.Println(nums)
	fmt.Println(media)
}
