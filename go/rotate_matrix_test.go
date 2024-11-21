package main

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(test *testing.T) {
	sums := rotateMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	expected := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}

	if !reflect.DeepEqual(sums, expected) {
		test.Fatalf("expected %v, got %v", expected, sums)
	}
}
