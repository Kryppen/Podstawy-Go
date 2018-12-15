package main

import (
	"fmt"
)

func Put100(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
		// if !ok {
		// 	fmt.Println("Error putting to a channel.")
		// 	return true
		// }

	}
	close(ch)
	//return false
}

func main() {
	channel := make(chan int)

	go func() {
		Put100(channel)
	}()

	for v := range channel {
		fmt.Println("Number pulled:", v)
	}
}
