package main

import (
	"fmt"
	"sync"
)

var balance int = 0
var mu sync.Mutex

func deposit(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	balance += 100
}

func main() {
	var wg sync.WaitGroup

	// Running 3 deposits
	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go deposit(&wg)
	}
	wg.Wait()
	fmt.Println("FInal balance: ₹", balance)
}
