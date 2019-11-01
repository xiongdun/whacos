package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 10)
	stringChan := make(chan string)
	//bootChan := make(chan bool)

	go channel1(intChan)
	go channel2(stringChan)
	//time.Sleep(time.Second * 6)

	for value := range intChan {

		fmt.Println("int channel`s value is ", value)

		//value1, ok1 := <-stringChan
		//if ok1 {
		//	fmt.Println("string channel`s value is ", value1)
		//}
	}

	//value1, ok1 := <-stringChan
	//if ok1 {
	//	fmt.Println("string channel`s value is ", value1)
	//}
	fmt.Println(<-stringChan)
	//<-intChan

}

func channel1(ch chan int) {
	for i := 0; i < 10; i++ {
		value := i + 1

		ch <- value
	}
	close(ch) // 使用完channel 一定要记得close
}

func channel2(ch chan string) {
	ch <- "my name is xiongdun"
	close(ch)
}
