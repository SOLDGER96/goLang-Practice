package main

import (
	"errors"
	"fmt"
	"os"
	//"os"
)

// Validate User Input
// Store Output in a file

func main() {
	revenue, err1 := getinput("Revenue: ")
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	expense, err2 := getinput("Expense: ")
	// if err2 != nil {
	// 	fmt.Println(err2)
	// 	return
	// }

	taxRate, err3 := getinput("Tax Rate: ")
	if err3 != nil || err2 != nil || err1 != nil {
		fmt.Println(err3)
		return
	}

	op1, op2, op3 := calc(revenue, expense, taxRate)
	disp(op1, op2, op3)
	storeResults(op1, op2, op3)
}

func getinput(ip string) (float64, error) {
	fmt.Print(ip)
	var value float64
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("Invalid Input")
	}

	return value, nil
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

func storeResults(ebt, profit, ratio float64) {
	output := fmt.Sprintf("EBT: %.1f \nProfit: %.1f \nRatio: %.2f",ebt, profit, ratio)
	os.WriteFile("output.txt", []byte(output), 0644)
}