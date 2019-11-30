package main

import (
	"fmt"
)

func main() {

	fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
}

func productExceptSelf(nums []int) []int {

	if len(nums) <= 0 {
		return nil
	}

	result := make([]int, len(nums))
	k := 1

	for i := 0; i < len(result); i++ {
		result[i] = k
		k = k * nums[i]
	}

	k = 1
	for i := len(result) - 1; i >= 0; i-- {
		result[i] *= k
		k *= nums[i]
	}

	return result

}
