package main

import (
	"fmt"
	"time"
)

func main() {
	go greet("Nice to Meet You!")
	go greet("Drte hai kya?")
	go slowGreet("it is slower!")
	greet("Took more time!")
}

func greet(str string) {
	fmt.Println("Hello! ", str)
}

func slowGreet(str string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello! ", str)
}
