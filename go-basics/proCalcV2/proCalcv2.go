package main

import (
	"fmt"
)

func main() {
	// var revenue float64
	// var expense float64
	// var taxrate float64


	// fmt.Print("Enter Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Enter Expense: ")
	// fmt.Scan(&expense)

	// fmt.Print("Enter TaxxRate: ")
	// fmt.Scan(&taxrate)

	revenue := getinput("Revenue: ")
	expense := getinput("Expense: ")
	taxRate := getinput("Tax Rate: ")

	// var earnB4Tax float64 = revenue - expense
	// var tax float64 = earnB4Tax * (taxrate/100)
	// var earnAfterTax float64 = earnB4Tax - tax
	// var ratio float64 = earnB4Tax / earnAfterTax

	// fmt.Println("Earning Before Tax: ", earnB4Tax)
	// fmt.Println("Earning After Tax: ", earnAfterTax)
	// fmt.Println("Ratio: ", ratio)

	op1, op2, op3 := calc(revenue, expense, taxRate)
	disp(op1, op2, op3)
}

func getinput(ip string) float64 {
	fmt.Print(ip)
	var value float64
	fmt.Scan(&value)
	return value
}

func calc(rev float64, exp float64, tax float64) (float64, float64, float64) {
	ebt := rev - exp
	profit := ebt * (1 - tax/100)
	ratio := ebt/profit
	return ebt, profit, ratio
}

func disp(op1,op2,op3 float64) {
	fmt.Printf("Earning b4 Tax: %.1f \nProfit: %.1f \nRatio: %.2f \n",op1, op2, op3)
}
