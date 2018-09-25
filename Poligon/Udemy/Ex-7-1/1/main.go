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

func half(i int) (wynik int, parzysta bool) {
	wynik = i / 2
	if i%2 == 0 {
		parzysta = true
	} else {
		parzysta = false
	}
	return wynik, parzysta
}
