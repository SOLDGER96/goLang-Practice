package main

import (
	"fmt"
	"sync"
)



var wg sync.WaitGroup

func main() {
	// - - - - - - Channel Syntax - - - - - -
	// mychannel := make(chan string)
	// mychannel <- "Secret Message"
	// receivedessage := <- mychannel
	wg.Add(1)
	go sayHello()
	fmt.Println("Hello from main!")
	wg .Wait()
}

func sayHello() {
	fmt.Println("Hello from the goroutine!")
	wg.Done()
}

// func main() {
// 	mychan := make(chan string)
// 	go sendData(mychan)
// 	recData := <- mychan
// 	fmt.Println(recData)
// }

// func sendData(c chan string) {
// 	c <- "Hello through the channel!"
// 

