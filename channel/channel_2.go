package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan int {
	c := make(chan int)
	//避免其他goroutine还未拿到channel 而进行打印造成deadlocak
	//真正做事情的worker 》
	go woker(id, c)
	return c
}

func woker(id int, c chan int) {
	for {
		//fmt.Printf("worker %d recived %c\n", id, <-c)

		//n, ok := <-c
		//if !ok {
		//	break
		//}
		//fmt.Printf("worker %d recived %c\n",
		//	id, n)

		for n := range c {
			fmt.Printf("worker %d recived %c\n",
				id, n)
		}
	}
}

//表明该worker只能发送数据
func createWorker2(id int) chan<- int {
	c := make(chan int)
	go woker(id, c)
	return c
}

//只能从channel里面接受数据
func createWorker3(id int) <-chan int {
	c := make(chan int)
	go woker(id, c)
	return c
}

//带缓冲的channel
func createWorker4(id int) chan<- int {
	c := make(chan int, 3)
	go woker(id, c)
	return c
}

func channelDemo2() {
	var chans [10]chan int
	for i := 0; i < 10; i++ {
		chans[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		chans[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		chans[i] <- 'A' + i
	}
}

func channelClosed() {
	// c:= make(chan int,3)
	// go woker(0,c)
	c := createWorker4(0)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	channelClosed()
}
