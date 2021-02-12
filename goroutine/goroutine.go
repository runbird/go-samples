package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(indx int) {
			//非抢占式协程，fmt.printf 磁盘I/O 切换主动交出使用权
			fmt.Printf("goroutine :%d\n", indx)
		}(i)
	}
	time.Sleep(time.Second)

	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			a[i]++
			//主动让出使用权
			runtime.Gosched()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(a)

	var b [10]int
	for i := 0; i < 10; i++ {
		//不定义go func里面的 i 会导致下标越界
		//使用 go run -race goroutine.go可查看原因
		//原因：闭包引用i,当执行到41行的时候，b[i] i=10,数组下标越界
		go func() {
			b[i]++
			runtime.Gosched()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}
