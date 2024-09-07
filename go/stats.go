package main

import "fmt"

func Sum(values []float64) float64 {
	var total float64

	for _, v := range values {
		total += v
	}

	return total
}

func Mean(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Mean of empty slice")
	}

	return Sum(values) / float64(len(values)), nil
}
