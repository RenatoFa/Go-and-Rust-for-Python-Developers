package main

import "fmt"

func chain() {
	ch := make(chan int)

	go func() {
		ch <- 99
	}()

	val := <-ch

	fmt.Printf("got %d\n", val)
}
