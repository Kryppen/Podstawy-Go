package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var Account, ReadK1, ReadK2 int64
var wg sync.WaitGroup

//var mut sync.Mutex

func Klient1() {
	for i := 0; i < 100; i++ {
		//mut.Lock()
		//runtime.Gosched()
		ReadK1 = atomic.LoadInt64(&Account)
		atomic.AddInt64(&Account, 1) //Account++
		fmt.Println("K1:", atomic.LoadInt64(&Account))
		//time.Sleep(10 * time.Millisecond)
		//mut.Unlock()
	}
	wg.Done()
}

func Klient2() {
	for i := 0; i < 100; i++ {
		//mut.Lock()
		//runtime.Gosched()
		ReadK2 = atomic.LoadInt64(&Account)
		atomic.AddInt64(&Account, 1) //Account++
		fmt.Println("K2:", atomic.LoadInt64(&Account))
		//time.Sleep(20 * time.Millisecond)
		//mut.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go Klient1()
	go Klient2()
	wg.Wait()
	fmt.Println(Account)

}
