package main

import "fmt"

func SecondBest(prices []int) (int, error) {
	if len(prices) == 0 {
		return 0, fmt.Errorf("SecondBest of empty prices")
	}

	first, second := prices[0], prices[0]

	for _, price := range prices[1:] {
		switch {
		case price > first:
			first, second = price, first
		case price > second:
			second = price
		}
	}

	return second, nil
}
