package main

import (
	"fmt"
)

func main() {
	var wyn int
	var parz bool
	wyn, parz = half(11)
	if parz {
		fmt.Println("Wynik=", wyn, " Liczba jest parzysta!")
	} else {
		fmt.Println("Wynik=", wyn, " Liczba jest nieparzysta!")

	}
}

func half(i int) (int, bool) {
	return i / 2, i%2 == 0
}
