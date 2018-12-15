package main

import (
	"fmt"
)

func send(c chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}

func main() {
	c := make(chan int)

	send(c)

	v := 0
	var ok bool
	for v := range c {
		fmt.Println(v, ok)

	}

	//close(c)

	v, ok = <-c
	fmt.Println(v, ok)

}
