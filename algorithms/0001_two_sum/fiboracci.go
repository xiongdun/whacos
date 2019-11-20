package main

import "fmt"

func main() {
	fmt.Println(fiboracci(3))

	fmt.Println(fibTail(3))

	fmt.Printf("%T", zero)
}

const zero = 0.0

func fiboracci(n int) int {
	if n < 2 {
		return n
	} else {
		return fiboracci(n-1) + fiboracci(n-2)
	}
}

func fibTail(n int) int {
	if n < 0 {
		return -1
	}
	if n < 3 {
		return 1
	}
	return fiboracci_tail(n, 1, 1, 3)
}

func fiboracci_tail(n, a, b, begin int) int {
	if n == begin {
		return a + b
	} else {
		return fiboracci_tail(n, b, a+b, begin+1)
	}
}
