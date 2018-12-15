package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mu sync.Mutex

func PutOnChannel(ch chan<- int) {
	for j := 0; j < 10; j++ {
		ch <- j
	}

}

func main() {

	channel := make(chan int)

	for i := 0; i < 10; i++ {
		go PutOnChannel(channel)
		fmt.Println("Threads:", runtime.NumGoroutine())

	}

	k := 0
	for {
		k++
		if k == 101 {
			return
		}

		v, ok := <-channel
		if !ok {
			fmt.Println("Exiting")
			return
		}
		fmt.Println("Number pulled:", v, ok, k)
	}

}
