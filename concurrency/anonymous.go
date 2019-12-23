package concurrency

import "fmt"

func main() {
	doneChan := make(chan string)
	//匿名 goroutine
	go func() {
		fmt.Println("runing in my own go routine")
	}()
	//阻塞到上面的goroutine给这个doenChan写入数据
	<-doneChan
}
