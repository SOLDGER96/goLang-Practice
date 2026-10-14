package main

import "fmt"

func sendData(c chan string) {
	c <- "This is go Routine!"
}

func main() {
	mychan := make(chan string)
	go sendData(mychan)
	recData := <-mychan
	fmt.Println(recData)
}
