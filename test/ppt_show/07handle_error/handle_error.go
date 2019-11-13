package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

// 之前在做其他语言是，对于异步调用，总有返回值和错误返回，并且能接收，对其进行处理。
//在go语言中，对于协程goroutine的代码，因为没有返回值接收，导致新手在学习时候会有困惑：
//goroutine的返回值（包括错误）怎么捕获呢？

// 首先我们定义一个结构体// 类似于class
type Resp struct {
	data  int
	error error
}

func main() {
	//handleError()

	go printNumber()

	time.Sleep(1000)

	channel1()

	runtime.GC()

}

func channel1() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go service1(ch1)
	go service1(ch2)

	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	case <-time.After(2 * time.Second):
		fmt.Println("no case ok!")
	}
}

func printNumber() {
	defer func() {
		fmt.Println("恢复程序运行")
		recover()
	}()

	go func() {
		fmt.Println("ssssssssssssss")
	}()

	panic("发生异常")
}

func handleError() {
	resp := make(chan Resp)
	stop := make(chan struct{})

	go func() {

		defer func() {
			panic("发生异常")
		}()

		t := time.Tick(time.Second)
		index := 0
		for {
			select {
			case <-t:
				resp <- Resp{
					data: index,
				}
				index++
			case <-stop:
				resp <- Resp{
					error: nil,
				}
			}
		}
	}()

	for {
		select {
		case value := <-resp:
			if value.error != nil {
				log.Fatal(value.error)
			}
			if value.data == 5 {
				stop <- struct{}{}

				panic("异常")
			}
			fmt.Println("index", value.data)
		}
	}
}

func service1(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "from service1"
}

func service2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from service2"
}
