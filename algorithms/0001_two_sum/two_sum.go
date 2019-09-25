package main

import "fmt"

func main() {
	sum := twoSum([]int{8, 2, 7, 11, 9, 10}, 9)

	fmt.Println(sum)
}

func twoSum(nums []int, target int) []int {

	ints := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := ints[another]; ok {
			return []int{ints[another], i}
		}
		ints[nums[i]] = i
	}
	return nums
}
