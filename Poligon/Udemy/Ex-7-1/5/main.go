package main

import (
	"fmt"
)

func main() {
	foo(1, 2, 3, 4)
	foo(1, 2, 3)
	foo()
	aSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	foo(aSlice...)
}

func foo(is ...int) {
	fmt.Println(is)
}
