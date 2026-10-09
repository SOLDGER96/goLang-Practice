package main

import "fmt"

func main() {
	prices := []int{1,2,3,4,5}
	fmt.Println(prices)

	prices = append(prices, 6,7,8)
	fmt.Println(prices)

	disPrices := []int{12,43,21}

	// Use the syntax below to append slice to a slice
	prices = append(prices, disPrices...)
}
