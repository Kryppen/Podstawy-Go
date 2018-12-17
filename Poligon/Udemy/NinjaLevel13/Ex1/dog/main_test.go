package dog

import (
	"fmt"
	"testing"

	"github.com/Kryppen/Podstawy-Go/Poligon/Udemy/NinjaLevel13/Ex1/dog"
)

//TestYears testuje wartość funkcji Years z pakietu dog
func TestYears(t *testing.T) {
	x := dog.Years(7)
	if x != 49 {
		t.Error("Powinniśmy mieć 49, a mamy", x)
	}
}

//TestYearsTwo testuje wartość funkcji Years z pakietu dog
func TestYearsTwo(t *testing.T) {
	x := dog.Years(7)
	if x != 49 {
		t.Error("Powinniśmy mieć 49, a mamy", x)
	}
}

//BenchamrkYears testuje prędkość działania funkcji Years z pakietu dog
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.Years(10)
	}
}

//BenchamrkYears testuje prędkość działania funkcji YearsTwo z pakietu dog
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.YearsTwo(10)
	}
}

//ExampleYears pokazuje działanie funkcji Years z pakietu dog.
func ExampleYears() {
	fmt.Println(dog.Years(10))
	//Output:
	//70
}

//ExampleYearsTwo pokazuje działanie funkcji YearsTwo z pakietu dog.
func ExampleYearsTwo() {
	fmt.Println(dog.YearsTwo(10))
	//Output:
	//70
}
