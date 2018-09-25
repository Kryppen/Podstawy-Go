package main

import (
	"fmt"
)

func main() {
	wyn := func(i int) (int, bool) {
		return i / 2, i%2 == 0
	}

	fmt.Println(wyn(11))
}
