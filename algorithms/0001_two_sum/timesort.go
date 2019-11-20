package main

import (
	"fmt"
	"time"
)

func main() {

	numbers := []int{43, 23, 21, 13, 33, 10, 14, 101, 98, 87, 51}

	ch := make(chan int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		go sleepSort(numbers[i], ch)
	}

	for i := 0; i < len(numbers); i++ {

		fmt.Printf("%d ", <-ch)
	}

	close(ch)
}

func sleepSort(number int, ch chan int) {
	time.Sleep(time.Millisecond * time.Duration(number))

	ch <- number
}
