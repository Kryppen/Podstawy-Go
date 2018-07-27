package main

import "fmt"
import "errors"
import "os"

const Pi = 3.1415

func main() {

	jedynka, dwojka, trojka := 1, 2, 3

	fmt.Printf("Cześć, świecie!\n")
	fmt.Printf("Jedynka to: %v, dwójka to: %v, trójka to: %v.\n", jedynka, dwojka, trojka)
	fmt.Printf("Liczba Pi to: %v\n", Pi)
	var i, j int32 = 10, 20
	var wynik int32 = i + j
	fmt.Printf("Wynik to: %v\n", wynik)

	napis := "Hello"
	tablica := []byte(napis)
	tablica[0] = 'Y'
	NapisZmieniony := string(tablica)
	fmt.Printf("Napis zmieniony to: %v", NapisZmieniony)
	KilkaLinii := `World 
	i następna linia`
	napis = napis + KilkaLinii
	fmt.Printf(napis)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("\n\n")
	fmt.Printf(os.Hostname())
}
