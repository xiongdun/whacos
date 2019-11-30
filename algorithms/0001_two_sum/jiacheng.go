package main

import "fmt"

func main() {
	fmt.Println(factorial2(5))
}

func factorial(n int) int {
	if n == 1 {
		return n
	} else {
		return n * factorial(n-1)
	}
}

func factorial2(n int) int {

	if n < 2 {
		return 1
	}

	return factorialTail(n, 1, 1, 2)
}

func factorialTail(n, a, b, begin int) int {
	if n == begin {
		return a * b
	} else {
		return factorialTail(n, b, a+b, begin+1)
	}
}
