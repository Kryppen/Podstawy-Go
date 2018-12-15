package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	ch := gen(q)

	receive(ch, q)

	fmt.Println("about to exit")
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			fmt.Println("Exit receive")
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 0
	}()

	return c
}
