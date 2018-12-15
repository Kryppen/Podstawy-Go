package main

import (
	"fmt"
	"math"
)

// func TriangleErea(a,h float32) float32 {
// 	return 1/2 * a * h
// }
//SquareArea counts area of a square
func SquareArea(a float32) float32 {
	return a * a
}

//CircleArea counc area of a circle
func CircleArea(r float32) float32 {
	return math.Pi * r * r
}

//Area couns any area
func Area(f func(x float32) float32, x float32) float32 {
	ret := f(x)
	return ret
}

func main() {
	sq := SquareArea
	area1 := Area(sq, 4)
	fmt.Println(area1)

	ca := CircleArea
	area2 := Area(ca, 4)
	fmt.Println(area2)
}
