package concurrency

import (
	"fmt"
	"time"
)

func channal() {
	//无缓冲channel
	//次只有一条数据适合通过 channel。
	//channel := make(chan string)
	//缓冲channel 容量为3的缓冲
	//缓冲 channel的工作原理与非缓冲 channel相似，只是有一点需要注意：
	// 我们可以在需要另外的 goroutine读取 channel之前将多条数据发送到 channel。
	//使用缓冲 channel不会阻止发生阻塞。例如，如果寻矿土拨鼠比采矿土拨鼠快10倍，
	// 并且他们通过大小为2的缓冲 channel进行通信，则寻矿土拨鼠仍将在程序中多次被阻塞。
	bufferChannel := make(chan string, 3)

	go func() {
		bufferChannel <- "first"
		fmt.Println("sent first")
		bufferChannel <- "second"
		fmt.Println("sent second")
		bufferChannel <- "third"
		fmt.Println("sent third")
	}()

	<-time.After(time.Second * 1)

	go func() {
		firstRead := <-bufferChannel
		fmt.Println("Receiving ...")
		fmt.Println(firstRead)
		secondRead := <-bufferChannel
		fmt.Println(secondRead)
		thirdRead := <-bufferChannel
		fmt.Println(thirdRead)
	}()
}
