package main

import (
	"fmt"
)

type intek int

var x intek
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x += 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)

}
