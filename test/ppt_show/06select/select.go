package main

import (
	"fmt"
	"strconv"
	"time"
)

func printStringNumber(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- "I am the sample one num :" + strconv.Itoa(i)
		//time.Sleep(3 * time.Second)
	}
	close(ch)
}

func main() {
	stringChan := make(chan string, 3)

	go printStringNumber(stringChan)

	for {

		select {
		case value, ok := <-stringChan:
			if !ok {
				fmt.Println("string channel is false")
				return
			}

			fmt.Println(value)
		case <-time.After(time.Second * 2):
			fmt.Println("time out ")
			//default:
			//	fmt.Println("default")
			//	return
		}
	}
	//time.Sleep(60 * time.Second)
}
