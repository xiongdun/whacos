package main

import (
	"fmt"
	"math"
)

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	i := search(num, 11)
	fmt.Println(i)

	fs := math.Log2(1000000000000)
	fmt.Println(fs)

	math.Pow()
}

func search(nums []int, item int) int {

	low := 0
	high := 0
	if nums[0] < nums[len(nums)-1] {
		low = 0
		high = len(nums) - 1
	} else {
		low = len(nums) - 1
		high = 0
	}

	for low <= high {
		mid := low + high
		guess := nums[mid]

		if guess == item {
			return mid
		} else if guess < item {
			low = mid + 1
		} else if guess > item {
			high = mid - 1
		}

	}
	return -1
}
