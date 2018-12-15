package main

import (
	"fmt"
	"math"
)

//Shape interface
type Shape interface {
	Area() float32
}

//Square struct
type Square struct {
	a float32
}

//Circle struct
type Circle struct {
	r float32
}

//Area func counts an area
func (s Square) Area() float32 {
	return s.a * s.a
}

//Area func counts an area
func (c Circle) Area() float32 {
	return math.Pi * c.r * c.r
}

//Info func prints an area
func Info(shape Shape) {
	fmt.Println(shape.Area())
}

func main() {

	Square1 := Square{
		a: 10.0,
	}
	Circle1 := Circle{
		r: 10.0,
	}
	Info(Square1)
	Info(Circle1)

}
