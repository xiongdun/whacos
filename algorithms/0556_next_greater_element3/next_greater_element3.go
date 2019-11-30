package main

import "fmt"

func main() {
	nextGreaterElement(1234)
}

func nextGreaterElement(n int) int {

	a := []byte(string(n))
	fmt.Println(a)

	i := len(a) - 2

	for i >= 0 && a[i+1] <= a[i] {
		i--
	}

	if i < 0 {
		return -1
	}

	j := len(a) - 1

	for j >= 0 && a[j] <= a[i] {
		j--
	}

	swap(a, i, j)
	reverse(a, i+1)

	return 0
}

func reverse(a []byte, start int) {
	i, j := start, len(a)-1

	for i < j {
		swap(a, i, j)
		i++
		j--
	}
}

func swap(a []byte, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
