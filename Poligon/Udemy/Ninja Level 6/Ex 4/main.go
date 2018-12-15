package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) Speak() {
	fmt.Println("Hello, I'm", p.first, "age", p.age)
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p.Speak()
}
