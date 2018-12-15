package main

import (
	"fmt"
)

type Person struct {
	Imie     string
	Nazwisko string
	Wiek     int
}

func (p *Person) Speak() {
	fmt.Println("Powiadam: ", p.Imie, p.Nazwisko, p.Wiek)
}

type Human interface {
	Speak()
}

func saySomething(h Human) {
	h.Speak()
}

func main() {
	osoba := Person{
		Imie:     "Jan",
		Nazwisko: "Kowalski",
		Wiek:     38,
	}

	saySomething(&osoba)
}
