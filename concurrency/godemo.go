package concurrency

import (
	"fmt"
	"time"
)

func godemo() {
	theMine := [5]string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)

	//Finder
	go func(mine [5]string) {
		for _, item := range mine {
			//发送数据
			oreChan <- item
		}
	}(theMine)

	//Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			//接收数据
			foundOre := <-oreChan
			fmt.Printf("Miner: Recevied {},from finder", foundOre)
		}
	}()
	<-time.After(time.Second * 5)
}
