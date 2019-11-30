package main

import "fmt"

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))

	prices2 := []int{2, 3, 5, 4, 8, 6, 6, 4}

	fmt.Println(maxProfit2(prices2))
}

// 动态规划，求最大收益
func maxProfit(prices []int) int {
	var miniProfit, maxProfit = int(^uint(0) >> 1), 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < miniProfit {
			miniProfit = prices[i]
		} else if prices[i]-miniProfit > maxProfit {
			maxProfit = prices[i] - miniProfit
		}
	}

	return maxProfit
}

// 贪心解法，只要是股票在涨就买，不涨就不买，求最优解
func maxProfit2(stocks []int) int {
	var profit = 0

	for i := 0; i < len(stocks); i++ {
		if i == 0 {
			continue
		}
		if stocks[i] > stocks[i-1] {
			temp := stocks[i] - stocks[i-1]
			profit += temp
		}
	}
	return profit
}
