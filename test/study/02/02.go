package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type user struct {
}

func (u user) getName() {
	fmt.Println("=============")

}

func main() {
	//x := 42
	//fmt.Println(x)

	//fmt.Println(runtime.NumCPU())

	//addCocurrent2 := addCocurrent
	//
	//addCocurrent2(2, []int{1, 2, 3, 4, 5, 6})

	org := []int{}
	fmt.Println(len(org), cap(org), org)
	bubbleSort(org)
	fmt.Println(len(org), cap(org), org)

	user{}.getName()

	fmt.Printf("%p\n", &org)
	//fmt.Printf("%p", &org[:])
	var a, b int
	a, org[0] = 1, 12
	fmt.Print(a, b, org[1])
}

//func find(topic string, docs []string) int {
//	var found int
//	for _, doc := range docs {
//		item , err := read()
//	}
//}
//
//func read(doc string) {
//	time.Sleep(10000)
//
//	var d goes.Document
//	fmt.Println(d)
//
//	if err := xml.Unmarshal([]byte(file)) {
//
//	}
//}

func bubbleSort(numbers []int) {
	n := len(numbers)

	for i := 0; i < n; i++ {
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []int, currentPass int) bool {
	var idx int
	idxNext := idx + 1
	n := len(numbers)
	var swap bool

	for idxNext < (n - currentPass) {
		a := numbers[idx]
		b := numbers[idxNext]

		if a > b {
			numbers[idx] = b
			numbers[idxNext] = a
			swap = true
		}
		idx++
		idxNext = idx + 1
	}
	return swap
}

func addCocurrent(goroutines int, numbers []int) int {
	var v int64
	totalNumber := len(numbers)

	lastGoroutines := goroutines - 1
	stride := totalNumber / goroutines

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride
			if g == lastGoroutines {
				end = totalNumber
			}
			var lv int
			for _, n := range numbers[start:end] {
				lv += n
			}
			atomic.AddInt64(&v, int64(lv))
			fmt.Println(lv)
			wg.Done()
		}(g)
	}
	wg.Wait()
	return int(v)
}
