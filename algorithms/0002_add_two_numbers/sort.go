package main

import (
	"fmt"
	"math/rand"
)

func main() {

	numbers := randomNumber()

	fmt.Println("==========排序前 ======")

	fmt.Println(numbers)
	fmt.Println("==========归并排序后 ======")
	//numbers = mergeSort(numbers)
	//fmt.Println(numbers)
	fmt.Println("==========快速排序后 ======")
	quickSort(numbers)
	fmt.Println(numbers)
	//selectSort(numbers)
	//fmt.Println(numbers)
	fmt.Printf("find value from numbers`s position is %v \n", find(numbers, 34))
}

// 生成随机数
func randomNumber() []int {

	numbers := make([]int, 0)

	for i := 0; i < 1000; i++ {
		numbers = append(numbers, rand.Intn(1500))
	}
	return numbers
}

// 二分查找法
func find(numbers []int, target int) int {

	left, right := 0, len(numbers)-1

	for left < right {

		mid := (left + right) / 2

		if numbers[mid] == target {
			return mid
		} else if numbers[mid] > target {
			right = mid - 1
		} else if numbers[mid] < target {
			left = mid + 1
		} else {
			return -1
		}
	}
	return -1
}

// 睡眠排序
func sleepSort(numbers []int) []int {

	return numbers
}

// 冒泡排序
func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {

		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

}

// 插入排序
func insertSort(numbers []int) {

	if len(numbers) < 1 {
		return
	}

	for i := 1; i < len(numbers); i++ {

		temp := numbers[i]
		for j := i - 1; j >= 0; j-- {

			if temp < numbers[j] {
				numbers[j+1] = numbers[j]
			}
			numbers[j+1] = temp
			break
		}
	}
}

// 选择排序
func selectSort(numbers []int) {

	for i := 0; i < len(numbers); i++ {

		for j := i; j < len(numbers); j++ {

			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

}

// 归并排序
func mergeSort(numbers []int) []int {

	//left, right := 0, len(numbers)

	if len(numbers) < 2 {
		return numbers
	}

	i := len(numbers) / 2

	left := mergeSort(numbers[0:i])
	right := mergeSort(numbers[i:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0

	l, r := len(left), len(right)

	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
			continue
		}
		result = append(result, left[m])
		m++
	}
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}

// 快速排序
func quickSort(numbers []int) {

	if len(numbers) < 1 {
		return
	}
	quick(numbers, 0, len(numbers)-1)
}

func quick(numbers []int, left, right int) {

	if left < right {
		//position := partition(numbers, left, right)
		//
		//partition(numbers, left, position-1)
		//partition(numbers, position+1, right)

		key := numbers[(left+right)/2]
		i, j := left, right

		for {
			for numbers[i] < key {
				i++
			}
			for numbers[j] > key {
				j--
			}

			if i >= j {
				break
			}

			swap(numbers, i, j)
		}

		quick(numbers, left, i-1)
		quick(numbers, j+1, right)
	}
}

func partition(numbers []int, left, right int) int {

	key := numbers[right]

	i := left - 1

	for j := left; j < right; j++ {
		if numbers[j] <= key {
			i++
			swap(numbers, i, j)
		}
	}

	swap(numbers, i+1, right)

	return i + 1
}

func swap(numbers []int, i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}
