package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange(coins []int, amount int) int {

	if amount <= 0 {
		return -1
	}

	return coinChanges(0, coins, amount)
}

const Max = int(^uint(0) >> 1)

func coinChanges(index int, coins []int, amount int) int {

	if index < len(coins) && amount > 0 {
		maxVal := amount / coins[index]

		minCost := Max

		for i := 0; i <= maxVal; i++ {
			if amount >= i*coins[index] {
				result := coinChanges(index+1, coins, amount-i*coins[index])

				if result != -1 {
					minCost = int(math.Min(float64(minCost), float64(result+i)))
				}
			}
		}
		if minCost == Max {
			minCost = -1
		}
		return minCost
	}
	return -1
}
