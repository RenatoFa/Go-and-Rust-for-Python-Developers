package main

import "fmt"

func fizzbuzz() {
	count, total := 0, 0
	for n := 1; n <= 100; n++ {
		if n%3 == 0 || n%5 == 0 {
			count++
			total += n
		}
	}
	fmt.Println(float64(total) / float64(count))
}
