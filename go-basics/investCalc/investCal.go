package main

import (
	"fmt"
	"math"
)
// Variables declared outside functions become global variables
const inflation = 6.0

func main() {
	//invAmount, expReturnRate, yrs := 1000.0, 7.5, 10.0
	
	
	var invAmount, yrs float64 = 1000, 10
	expReturnRate := 7.5
	//var yrs float64 = 10

	// constant cannot be used to assign to a PTR
	//fmt.Scan(&inflation) --> Shows error

	//fmt.Scan(invAmount)  --> It is valid coz we used a variable to assign to a PTR
	fmt.Print("Investment Amount: ")
	fmt.Scan(&invAmount)

	fmt.Print("Years: ")
	fmt.Scan(&yrs)

	fmt.Print("Expected Rate of Return: ")
	fmt.Scan(&expReturnRate)

	futureVal := invAmount * math.Pow(1 + expReturnRate / 100, yrs)
	futureRealVal := futureVal / math.Pow(1+inflation/100, yrs)
	fmt.Printf("Total Value: %f\n", futureVal)
	fmt.Printf("Total Value: %.2f\n", futureVal)
	fmt.Println(futureRealVal)


	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureVal)
	fmt.Print(formattedFV)

	// Adding Line Break in print statement
	fmt.Printf(`Future Value: %.1f\n 
	Future Value (Inflation): %.1f\n`, futureVal, futureRealVal)


	// Using a custom function
	opText("Tanmay")
}
func opText(t string) {
	fmt.Print(t)
}
