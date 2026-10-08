package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expense float64
	taxrate := 10.0

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expense: ")
	fmt.Scan(&expense)

	var earnB4Tax float64 = revenue - expense
	var tax float64 = earnB4Tax * (taxrate/100)
	var earnAfterTax float64 = earnB4Tax - tax
	var ratio float64 = earnB4Tax / earnAfterTax

	fmt.Println("Earning Before Tax: ", earnB4Tax)
	fmt.Println("Earning After Tax: ", earnAfterTax)
	fmt.Println("Ratio: ", ratio)
}