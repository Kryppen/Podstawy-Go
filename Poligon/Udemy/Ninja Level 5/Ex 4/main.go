package main

import (
	"fmt"
)

func main() {

	persona := struct {
		first string
		last  string
	}{
		"Adam",
		"Abacki",
	}

	fmt.Println(persona)

}
