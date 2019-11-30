package main

import "fmt"

func main() {

	nums := []int{2, 5, 5, 11}

	target := 10

	sum2 := two_sum2(nums, target)
	fmt.Println(nums[sum2[0]], nums[sum2[1]])

	for _, v := range sum2 {

		fmt.Printf("%d ", nums[v])
	}

}

func two_sum2(nums []int, target int) []int {

	m := make(map[int]int, len(nums))
	res := []int{}
	for k, v := range nums {
		if _, val2 := m[target-v]; val2 {
			res = []int{k, m[target-v]}
			break
		}
		m[v] = k
	}
	return res
}

func two_sum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Printf("因为nums[%d] + nums[%d] = %d + %d = %d\n", i, j, nums[i], nums[j], target)

				return []int{i, j}
			}
		}
	}
	return nil
}
