package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(tanxin(9))
}

func tanxin(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}

	if length == 3 {
		return 2
	}

	timesOf3 := length / 3

	if length-timesOf3*3 == 1 {
		timesOf3 -= 1
	}

	timesOf2 := (length - timesOf3*3) / 2

	return int(math.Pow(float64(3), float64(timesOf3)) * math.Pow(float64(2), float64(timesOf2)))
}
