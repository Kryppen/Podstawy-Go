package main

import (
	"fmt"
)

type vehicle struct {
	doors  int
	colors string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck1 := truck{
		vehicle: vehicle{
			doors:  2,
			colors: "white",
		},
		fourWheel: true,
	}
	sedan1 := sedan{
		vehicle: vehicle{
			doors:  4,
			colors: "black",
		},
		luxury: true,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1)

	fmt.Println(truck1.doors, truck1.colors, truck1.fourWheel)
	fmt.Println(sedan1.doors, sedan1.colors, sedan1.luxury)
}
