package main

import "fmt"

func main() {
	var suma int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			suma += i
		} else if i%5 == 0 {
			suma += i
		}

	}
	fmt.Println("Suma: ", suma)
}
