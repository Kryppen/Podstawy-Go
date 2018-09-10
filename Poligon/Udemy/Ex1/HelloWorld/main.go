package main

import (
	"fmt"
)

func main() {
	var Imie string
	fmt.Print("Podaj imie: ")
	fmt.Scanln(&Imie)
	fmt.Println("Witaj, ", Imie)
}
