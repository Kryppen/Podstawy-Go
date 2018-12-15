package main

import (
	"fmt"
)

func main() {
	type person struct {
		FirstName string
		LastName  string
		Flavors   []string
	}

	p1 := person{
		FirstName: "Adam",
		LastName:  "Abacki",
		Flavors:   []string{"Strawberry", "blueberry"},
	}

	p2 := person{
		FirstName: "Bartosz",
		LastName:  "Babacki",
		Flavors:   []string{"Blackberry", "Orange"},
	}

	persons := map[string]person{
		p1.LastName: p1,
		p2.LastName: p2,
	}

	for i, pers := range persons {
		fmt.Println("Person ", i, ": ", pers.FirstName, " ", pers.LastName)
		for _, fl := range pers.Flavors {
			fmt.Println(fl, "\t")
		}

	}

}
