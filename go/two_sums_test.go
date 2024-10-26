package main

import (
	"reflect"
	"testing"
)

func TestTwoSums(test *testing.T) {
	sums := TwoSums([]int{2, 7, 11, 15}, 9)

	expected := []int{1, 2}

	if !reflect.DeepEqual(sums, expected) {
		test.Fatalf("expected %v, got %v", expected, sums)
	}
}
