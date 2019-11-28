package concurrency

import (
	"fmt"
	"time"
)

func gochannel() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)

	//Finder
	go func(mine [5]string) {
		for _, item := range theMine {
			if item == "ore" {
				//发送数据给oreChannel
				oreChannel <- item
			}
		}
	}(theMine)

	//ore Breaker
	//可以使用 range 遍历
	go func() {
		for i := 0; i < 3; i++ {
			//从 oreChannel读取数据
			foundOre := <-oreChannel
			fmt.Println("From Finder: ", foundOre)
			//推送数据给minedOreChan
			minedOreChan <- "minedOre"
		}
	}()

	//Smelter
	go func() {
		for i := 0; i < 3; i++ {
			//从 minedOreChan 读取数据
			minedOre := <-minedOreChan
			fmt.Println("Find Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()

	//主函数定时器，防止主函数运行结束退出
	<-time.After(time.Second * 5)

	//range channel
	//一个 channel上使用 range将被阻塞，直到在 channel上发送了另一个数据。在所有需要发送的数据发送完后，
	//停止 goroutine被阻塞的唯一方法是用“close(channel)”关闭 channel
	go func() {
		for founder := range oreChannel {
			fmt.Println("Miner: Received " + founder + "from finder")
		}
	}()
}
