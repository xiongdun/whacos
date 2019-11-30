package main

import "fmt"

func main() {
	fmt.Println(movingCount(18, 35, 37))
}

func movingCount(threshold, rows, cols int) int {
	if threshold < 0 || rows < 0 || cols <= 0 {
		return 0
	}

	visited := make([]bool, rows*cols+1)

	for i := 1; i <= rows*cols; i++ {
		visited[i] = false
	}

	return movingCountCore(threshold, rows, cols, 0, 0, visited)
}

func movingCountCore(threshold, rows, cols, row, col int, visited []bool) int {
	count := 0
	if check(threshold, rows, cols, row, col, visited) {
		visited[row*cols+col] = true

		count = 1 + movingCountCore(threshold, rows, cols, row-1, col, visited) + movingCountCore(threshold, rows, cols, row, col-1, visited) + movingCountCore(threshold, rows, cols, row+1, col, visited) + movingCountCore(threshold, rows, cols, row, col+1, visited)
	}
	return count
}

func check(threshold, rows, cols, row, col int, visited []bool) bool {
	if row > 0 && row < rows && col >= 0 && col < cols && getDigitSum(row)+getDigitSum(col) <= threshold && !visited[row*cols+col] {
		return true
	}
	return false
}

func getDigitSum(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}

	return sum
}
