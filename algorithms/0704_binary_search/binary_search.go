package main

import "fmt"

func main() {

	nums := []int{-1, 0, 3, 5, 9, 12}

	target := 2

	fmt.Println(search(nums, target))

	fmt.Println()
}

func search(nums []int, target int) int {

	if len(nums) <= 0 {
		return -1
	}

	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid

		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return -1
		}
	}

	return -1
}
