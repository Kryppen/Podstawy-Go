package main

import (
	"fmt"
	"math"
)

func main() {
	x := func(r float32) float32 {
		return math.Pi * r * r
	}

	fmt.Println(x(5))
}
