package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(allCellsDistOrder(10, 10, 0, 0))
}

func allCellsDistOrder(R, C, r0, c0 int) [][]int {

	distance := make(map[int][][]int)

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {

			dis := math.Abs(float64(i-r0)) + math.Abs(float64(j-c0))

			if _, ok := distance[int(dis)]; ok {
				distance[int(dis)] = [][]int{{i, j}}
			} else {
				distance[int(dis)] = nil
				distance[int(dis)] = [][]int{{i, j}}
			}
		}
	}

	result := [][]int{}

	fmt.Println(distance)

	return result
}
