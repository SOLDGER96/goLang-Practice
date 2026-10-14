package main

import "fmt"

func ExecuteTask(callBack func()) {
	callBack()
}

func callBack() {
	fmt.Println("Just a callback function!")
}

func main() {
	ExecuteTask(callBack)
}