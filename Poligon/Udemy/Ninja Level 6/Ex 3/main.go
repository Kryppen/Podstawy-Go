package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()

}

func foo() {
	defer func() {
		fmt.Println("defered func in foo ran")
	}()
	fmt.Println("foo ran")
}

func bar() {
	fmt.Println("bar ran")
}
