package main

import (
	"fmt"
	"strconv"
)

func main() {

	defer func() {
		recover()
	}()

	defer f1()

	panic("2")
}

func f1() {
	fmt.Println("ssss")
}

func two(n int) {

	ints := [1]int{}
	fmt.Println(ints)

	if n <= 0 {
		return
	}

	s := ""

	for i := 0; i < n; i++ {
		s += "9"
	}

	fmt.Println(s)

	r, err := strconv.Atoi(s)

	if err != nil {
		return
	}

	for i := 0; i <= r; i++ {
		fmt.Println(i)
	}

}
