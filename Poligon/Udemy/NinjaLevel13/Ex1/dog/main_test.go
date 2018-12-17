package dog

import (
	"fmt"
	"testing"
)

//TestYears testuje wartość funkcji Years z pakietu dog
func TestYears(t *testing.T) {
	x := Years(7)
	if x != 49 {
		t.Error("Powinniśmy mieć 49, a mamy", x)
	}
}

//TestYearsTwo testuje wartość funkcji Years z pakietu dog
func TestYearsTwo(t *testing.T) {
	x := Years(7)
	if x != 49 {
		t.Error("Powinniśmy mieć 49, a mamy", x)
	}
}

//BenchamrkYears testuje prędkość działania funkcji Years z pakietu dog
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

//BenchamrkYears testuje prędkość działania funkcji YearsTwo z pakietu dog
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

//ExampleYears pokazuje działanie funkcji Years z pakietu dog.
func ExampleYears() {
	fmt.Println(Years(10))
	//Output:
	//70
}

//ExampleYearsTwo pokazuje działanie funkcji YearsTwo z pakietu dog.
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//Output:
	//70
}
