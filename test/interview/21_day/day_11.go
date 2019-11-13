package main

import "fmt"

func main() {

	f := day31
	f()

	ints := make([]int, 5)

	is := [5]int{1, 2, 34, 4}

	fmt.Printf("%p\n", &ints)
	fmt.Printf("%p\n", &is)
	day311(is[:]...)

	fmt.Println(is)
	fmt.Printf("%p", &is)
}

func day311(s ...int) {
	s1 := append(s, 4, 5, 6)
	fmt.Println(s1)
	fmt.Printf("%p , %p", &s, &s1)
}

func day31() {
	var slice []int = nil
	fmt.Println(slice, len(slice), cap(slice))
}

func day30() {
	slice := make([]int, 5)

	slice[0] = 1
	slice[1] = 2

	change(slice...)

	fmt.Println(slice)

	change(slice[:2]...)
	fmt.Println(slice)

	// 可变参数可以不用传值，默认为nil
	//change()
}

func change(s ...int) {
	s = append(s, 3)
	fmt.Println(s)

	//s[2] = 3
	//fmt.Println(len(s), cap(s))
	//fmt.Println(&s)
}
