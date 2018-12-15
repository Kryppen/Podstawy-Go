package main

import (
	"fmt"
)

func main() {
	skladniki := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(skladniki...))
	fmt.Println(bar(skladniki))
}

func foo(ii ...int) int {

	sum := 0
	for _, i := range ii {
		sum += i
	}
	return sum

}

func bar(ii []int) int {
	sum := 0
	for _, i := range ii {
		sum += i
	}
	return sum

}
