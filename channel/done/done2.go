package main

import (
	"fmt"
	"sync"
)

type worker2 struct {
	in chan int
	// wg *sync.WaitGroup
	done func()
}

func createWorker2(i int, wg *sync.WaitGroup) worker2 {
	w := worker2{
		in: make(chan int),
		//wg: wg,
		done: func() {
			wg.Done()
		},
	}
	//go doWork2(i,w.in,wg)
	go doWork2(i, w)
	return w
}

func doWork2(id int, w worker2) {
	for n := range w.in {
		fmt.Printf("worker %d recived %c\n", id, n)
		//wg.Done()
		w.done()
	}
}

func doneDemoe2() {
	var workers [10]worker2
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i, &wg)
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	doneDemoe2()
}
