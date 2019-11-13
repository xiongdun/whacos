package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go printNumber(1, 3, ch)
	go printNumber(4, 6, ch)

	_ = <-ch
	_ = <-ch
}

func printNumber(from, to int, ch chan int) {
	for i := from; i < to; i++ {
		fmt.Printf("%d\n", i)
		time.Sleep(1 * time.Second)
	}
	ch <- 0
}
