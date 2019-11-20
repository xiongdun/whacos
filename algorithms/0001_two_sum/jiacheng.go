package main

import "fmt"

func main() {
	fmt.Println(factorial(2))
}

func factorial(n int) int {
	if n == 1 {
		return n
	} else {
		return n * factorial(n-1)
	}
}
