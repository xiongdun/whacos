package main

import "fmt"

func main() {

	//1.给定一个整数数组=[6，4，-3，5，-2，-1，0，1，-9]，
	// 用任何语言实现一个函数，将所有正数向左移动，将所有负数向右移动。
	// 尽量使其时间复杂度为O（n），空间复杂度为O（1）。
	nums := [10]int{6, 4, -3, 5, -2, -1, 1, 0, -9, 10}
	fmt.Println(nums)

	// 从左到右遍历
	for i := len(nums); i > 0; i-- {

		// 从右到左遍历
		for j := len(nums) - 1; j > 0; j-- {
			// 如果左边的数值小于零
			// 那就和右边的大于零的值交换位置
			//if nums[j] >= 0 {
			if nums[j] > nums[j-1] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}

		}
	}
	fmt.Println()

	fmt.Println(nums)
}
