package main

import (
	"fmt"
	"testing"
)

func TestSecondBest(test *testing.T) {
	prices := []int{1, 0, 2, 3}

	p, err := SecondBest(prices)

	if err != nil {
		test.Fatal(err)
	}

	expected := 2

	if p != expected {
		test.Fatalf("expected %d, got %d", p, expected)
	}
}

var secondBestCases = []struct {
	prices   []int
	expected int
}{
	{[]int{1}, 1},
	{[]int{1, 2, 3, 4, 2}, 3},
}

func TestSecondBestTablet(test *testing.T) {
	for _, tc := range secondBestCases {
		name := fmt.Sprintf("%v", tc.prices)
		test.Run(name, func(t *testing.T) {
			p, err := SecondBest(tc.prices)
			if err != nil {
				test.Fatal(err)
			}
			if p != tc.expected {
				test.Fatalf("expected %d, got %d", p, tc.expected)
			}
		})
	}
}

func ExampleSecondBest() {
	values := []int{2, 1, 3}
	p, err := SecondBest(values)

	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println(p)
}
