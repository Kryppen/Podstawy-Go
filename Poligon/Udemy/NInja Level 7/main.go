package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

//ChangeMe changes person
func ChangeMe(p *person) {
	// (*p).first = "Adam"
	// (*p).last = "Abacki"
	// (*p).age = 40

	p.first = "Adam"
	p.last = "Abacki"
	p.age = 40
}

func main() {
	p1 := person{
		first: "Jan",
		last:  "Kowalski",
		age:   30,
	}

	fmt.Println("Before", p1)
	ChangeMe(&p1)
	fmt.Println("After", p1)

}
