package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(i int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(i, w.in, w.done)
	return w
}

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d recived %c\n", id, n)
		go func() { done <- true }()
	}
}

func doneDemoe() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, woker := range workers {
		woker.in <- 'a' + i
	}

	for i, woker := range workers {
		woker.in <- 'A' + i
	}
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	doneDemoe()
}
