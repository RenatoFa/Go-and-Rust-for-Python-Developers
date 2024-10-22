package main

type Addable interface {
	int | float64 | string
}

func Add[T int | float64 | string](a, b T) T {
	return a + b
}

func AddTwo[T Addable](a, b T) T {
	return a + b
}
