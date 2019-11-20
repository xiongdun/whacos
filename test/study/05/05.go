package main

import (
	"fmt"
	"runtime"
)

func main() {
	day05()
}

func func1() (int, int) {
	return 1, 2
}

func day05() {
	var x int
	x, y := func1()

	fmt.Println(x, y)
}

func day04() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	//for {}
}

var x = 23

func day03() {
	x := 2*x - 4
	println(x)
}

var f = func(int2 int) {
	print("x")
}

func day02() {
	f := func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}

	f(10)
}

func day01() {
	funs := test()

	for _, f := range funs {
		f()
	}
}

func test() []func() {
	var funs []func()

	var j int
	for i := 0; i < 2; i++ {
		j = i
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	fmt.Println(j)
	return funs
}
