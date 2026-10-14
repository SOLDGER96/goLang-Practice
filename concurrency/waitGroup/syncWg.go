package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from goRoutine!")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello(&wg)
	fmt.Println("Hello from Main")
	wg.Wait()
}