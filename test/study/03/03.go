package main

import "fmt"

type T struct {
	n int
}

func main() {
	func04()

}

func func05() {

}

func func04() {
	ints := make(map[string]int)

	ints["st"] = 1
	delete(ints, "st")
	fmt.Println(ints)
}

func func03() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", slice)
	c := slice[:2] // 0<= e <2
	fmt.Println(c)
	fmt.Printf("%p\n", c)
	slice[1] = 5
	fmt.Println(slice)
	fmt.Println(c)

	b := slice[:2:2]
	b = append(b, 7)
	fmt.Printf("%p\n", b)
	fmt.Println(b)
}

func func02() {

	slice1 := make([]int, 0, 5)

	fmt.Printf("%p\n", slice1)
	slice1 = append(slice1, 1, 2, 3)

	fmt.Printf("%p\n", slice1)
	slice1 = append(slice1, 4, 5)
	fmt.Printf("%p\n", slice1)
	slice1 = append(slice1, 6)
	fmt.Printf("%p\n", slice1)
	fmt.Println(len(slice1), cap(slice1))
}

func func01() {
	ts := [2]T{}

	for i, t := range &ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
			ts[0].n = 1
		case 1:
			fmt.Println(t.n, " ")
		}
	}
	fmt.Println(ts)

}
