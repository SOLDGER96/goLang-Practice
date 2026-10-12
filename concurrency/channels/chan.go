package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 4)
	
	go greet("Nice to Meet You!", done)
	go greet("Drte hai kya?", done)
	go slowGreet("it is slower!", done)
	go greet("Took more time!", done)
	<-done
	<-done
	<-done
	<-done

	fmt.Println("----------------------------")
	// Another approach 
	dones := make([]chan bool, 4)
	
	dones[0] = make(chan bool)
	go greet("Nice to Meet You!", dones[0])

	dones[1] = make(chan bool)
	go greet("Drte hai kya?", dones[1])

	dones[2] = make(chan bool)
	go slowGreet("it is slower!", dones[2])

	dones[3] = make(chan bool)
	go greet("Took more time!", dones[3])

	for _, done := range dones {
		<- done
	}

	fmt.Println("----------------------------")

	done2 := make(chan bool, 4)
	
	go greet("Nice to Meet You!", done2)
	go greet("Drte hai kya?", done2)
	go slowGreet("it is slower!", done2)
	go greet("Took more time!", done2)
	<-done2
	<-done2
	<-done2
	<-done2
	for range done {
		//fmt.Println(doneChan)
	}
}

func greet(str string, doneChan chan bool) {
	fmt.Println("Hello! ", str)
	doneChan <- true
}

func slowGreet(str string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello! ", str)
	doneChan <- true
	close(doneChan)
}
