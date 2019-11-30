package main

import "fmt"

func main() {

	ints := []int{6, 4, -3, 5, -2, -1, 0, 1, -9}

	swap(ints)

	fmt.Println(ints)

}

func swap(numbers []int) {

	if len(numbers) < 1 {
		return
	}

	left, right := 0, len(numbers)-1

	for left < right {

		for left < right && numbers[left] > 0 {
			left++
		}

		for left < right && numbers[right] < 0 {
			right--
		}

		if left < right {
			numbers[right], numbers[left] = numbers[left], numbers[right]
		}

	}
}
