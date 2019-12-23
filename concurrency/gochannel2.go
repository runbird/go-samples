package concurrency

import (
	"fmt"
	"time"
)

func main() {
	//channal上非阻塞式读取
	mychan := make(chan string)

	go func() {
		mychan <- "Message!"
	}()

	select {
	case msg := <-mychan:
		fmt.Println(msg)
	default:
		fmt.Println("No Msg")
	}

	<-time.After(time.Second * 1)

	select {
	case msg := <-mychan:
		fmt.Println(msg)
	default:
		fmt.Println("No Msg")
	}

	//在channel上非阻塞式写
	select {
	case mychan <- "message":
		fmt.Println("GET Message")
	default:
		fmt.Println("No Message")
	}
}
