package main

import (
	"reflect"
	"testing"
)

func TestThreeSums(test *testing.T) {
	sums := ThreeSums([]int{-1, 0, 1, 2, -1, -4})

	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	if !reflect.DeepEqual(sums, expected) {
		test.Fatalf("expected %v, got %v", expected, sums)
	}
}
