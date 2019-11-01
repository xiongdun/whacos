package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	// 无返回值协程
	go printHello()

	wg.Add(1)
	// 有返回值协程，但无法接收
	go returnValue()

	wg.Add(1)
	// 匿名协程
	go func() {
		defer wg.Done()
		fmt.Println(" 这里是匿名函数！")
	}()

	wg.Wait()
	// 主协程打印方法
	fmt.Println(" main function!")
}

func printHello() {
	defer wg.Done()
	fmt.Println("Hello World!")
}

func returnValue() string {
	defer wg.Done()
	return "this is the return string value!"
}
