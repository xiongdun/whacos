package main

import (
	"fmt"
	"time"
)

func foo() *int {
	var x int
	return &x
}

func bar() int {
	x := new(int)
	*x = 1
	return *x
}

func main() {

	//defer recover()
	//panic(1)
	//bar()

	print2()

}

func print1() {
	defer func() {
		i := recover()
		fmt.Println(i)
		fmt.Print(i)
		//fmt.Println("00000000000")
	}()

	defer func() {
		i := recover()
		fmt.Println(i)
		defer fmt.Print(i)
		panic(1)
	}()

	//defer recover()
	panic(2)
}

func print2() {
	defer func() {
		defer func() {
			recover()
		}()
	}()
	panic(1)
}

type span struct {
	time.Time
}
