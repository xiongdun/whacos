package main

import "fmt"

func main() {
	// 无返回值协程
	go printHello()

	//printHello()

	// 有返回值协程，但无法接收
	go returnValue()

	// 匿名协程
	go func() {
		fmt.Println(" 这里是匿名函数！")
	}()

	// 主协程打印方法
	fmt.Println(" main function!")
}

func printHello() {
	fmt.Println("Hello World!")
}

func returnValue() string {
	return "this is the return string value!"
}
