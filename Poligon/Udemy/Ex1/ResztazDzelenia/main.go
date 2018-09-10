package main

import "fmt"

func main() {
	var Dzielna, Dzielnik int32
	fmt.Println("Podaj dwie liczby: ")
	fmt.Scan(&Dzielna)
	fmt.Scan(&Dzielnik)

	if Dzielna < Dzielnik {
		fmt.Println("Błąd: Pierwsz licznba powinna być większa lub równa drugiej!")
	} else {
		var Wynik int32
		Wynik = Dzielna % Dzielnik
		fmt.Printf("Reszta z dzielenia to: %v", Wynik)
	}
}
