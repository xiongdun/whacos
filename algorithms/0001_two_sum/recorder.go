package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//numbers := []int{2, 4, 6, 8, 1, 3, 5, 7, 9}
	//numbers := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}

	fmt.Println("================操作前===============")

	fmt.Println(numbers)
	fmt.Println("================操作后===============")
	fmt.Println(recorder(numbers))
}

func recorder(numbers []int) []int {

	start := 0
	end := len(numbers) - 1

	for start < end {

		for start < end && numbers[start]&1 == 1 {
			start++
		}

		if start < end && numbers[end]&1 == 0 {
			end--
		}

		if start != end {
			temp := numbers[start]
			numbers[start] = numbers[end]
			numbers[end] = temp
		}
	}

	return numbers
}
