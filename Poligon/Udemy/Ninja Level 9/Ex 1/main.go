package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func Liczykrupa() {
	for i := 0; i < 100; i++ {
		fmt.Println("Liczykrupa:", i)
		//time.Sleep(250 * time.Millisecond)
	}
	wg.Done()
}

func Fibonacchi() {
	i := 1
	fibo := 1
	poprz := 1
	aktualny := 1
	fmt.Println("Fibonacchi:", i)
	fmt.Println("Fibonacchi:", i)
	for i < 100 {
		fibo = aktualny + poprz
		poprz = aktualny
		aktualny = fibo
		fmt.Println("Fibonacchi:", i, fibo)
		//time.Sleep(250 * time.Millisecond)
		i++
	}
	wg.Done()
}

func main() {
	fmt.Println("CPU start:", runtime.NumCPU())
	fmt.Println("Go threads start:", runtime.NumGoroutine())
	wg.Add(2)
	go Liczykrupa()
	go Fibonacchi()
	fmt.Println("CPU during:", runtime.NumCPU())
	fmt.Println("Go threads during:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("CPU end:", runtime.NumCPU())
	fmt.Println("Go threads end:", runtime.NumGoroutine())

}
