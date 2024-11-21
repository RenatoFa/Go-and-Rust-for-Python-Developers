package main

func rotateMatrix(matrix [][]int) [][]int {
	left, right := 0, len(matrix)-1

	for left < right {
		for index := range right - left {
			top, bottom := left, right

			topLeft := matrix[top][left+index]

			matrix[top][left+index] = matrix[bottom-index][left]
			matrix[bottom-index][left] = matrix[bottom][right-index]
			matrix[bottom][right-index] = matrix[top+index][right]
			matrix[top+index][right] = topLeft
		}

		right -= 1
		left += 1
	}

	return matrix
}
