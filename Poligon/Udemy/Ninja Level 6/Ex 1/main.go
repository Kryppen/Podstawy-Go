package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() float32 {
	return math.Pi
}

func bar() (int, string) {
	return 1, "To be or not to be?"
}
