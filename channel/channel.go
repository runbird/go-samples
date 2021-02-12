package main

import (
	"fmt"
	"time"
)

func channleDemo() {
	var chans [10]chan int
	for i := 0; i < 10; i++ {
		chans[i] = make(chan int)
		go worker(i, chans[i])
	}
	for i := 0; i < 10; i++ {
		chans[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		chans[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func worker(id int, c chan int) {
	for {
		fmt.Printf("worker %d recived %c\n",
			id, <-c)
	}
}

func main() {
	channleDemo()
}
