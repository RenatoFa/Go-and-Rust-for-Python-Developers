package main

import (
	"fmt"
	"time"
)

func worker(n int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("goroutine %d\n", n)
}
