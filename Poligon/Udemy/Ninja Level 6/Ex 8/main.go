package main

import (
	"fmt"
)

//TriangleArea calc func
func TriangleArea(a, h float32) func() float32 {
	return func() float32 {
		return 1 / 2 * a * h
	}
}

func main() {
	f := TriangleArea(2, 2)
	fmt.Printf("returned f type %T\n", f)
	a := f()
	fmt.Printf("returned area %T; area = %v\n", a, a)
}
