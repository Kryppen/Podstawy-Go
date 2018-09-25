package main

import (
	"fmt"
)

func main() {
	fmt.Println(Najwieksza(7, 2, 3, 4, 5, 6))
}

//Najwieksza szuka największej wartościw ciągu
func Najwieksza(is ...int) (max int) {
	for _, i := range is {
		if max < i {
			max = i
		}
	}
	return max
}
