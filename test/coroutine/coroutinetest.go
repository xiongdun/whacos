package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type Point struct {
}

func f(from string) {

	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	fmt.Println("====================================")
	f("direct")

	//go f("goroutine")
	//
	//go func(msg string) {
	//	fmt.Println(msg)
	//}("going")

	//time.Sleep(time.Second)
	fmt.Println("done")

	fmt.Println("=======================================")

	//go foo()
	//<-ch

	//fmt.Println("===========================")

	//go loop()

	//<-complete

	//go say("xiongdun")
	//<-ch1

	c, quit := make(chan int, 1), make(chan int)

	go func() {
		c <- 1    // c通道的数据没有被其他goroutine读取走，堵塞当前goroutine
		quit <- 0 // quit始终没有办法写入数据
	}()

	//<- c
	<-quit // quit 等待数据的写

	err := errors.New("this is the original error!")
	wrap := fmt.Errorf("warpping err be %w", err)
	fmt.Println(wrap)
}

var ch = make(chan int)
var ch1 = make(chan int)
var ch2 = make(chan int)
var complete = make(chan int)

func say(s string) {
	fmt.Println(s)
	ch1 <- <-ch2
}

func foo() {
	ch <- 0
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}

	complete <- 0
}
