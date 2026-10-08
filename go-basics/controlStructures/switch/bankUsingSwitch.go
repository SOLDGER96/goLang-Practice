package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const balanceFile string = "balance.txt"

func main() {

	var accBal, err = getBalFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		return
		//panic("Can't continue, No account found!")
	}


	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Account Balance: ", accBal)
		case 2: 
			fmt.Print("Your Deposit Amount: ")
			var dep float64
			fmt.Scan(&dep)
			if dep <= 0 {
				fmt.Print("Invalid Amount, Must be greater then 0")
				//return
				continue 
			}
			accBal += dep
			fmt.Println("Your Balance: ", accBal)
			writeToafile(accBal)
		case 3:
			fmt.Print("Enter Withdrawl Amount: ")
			var w float64
			fmt.Scan(&w)
			if w <= 0 {
				fmt.Print("Invalid Amount, Must be greater then 0")
				//return
				continue
			}
			if w > accBal {
				fmt.Println("Invalid Amount, Not enough Balance!")
				//return
				continue
			}
			accBal -= w
			fmt.Println("Your Balance: ", accBal)
			writeToafile(accBal)
		default:
			fmt.Println("Thank you!")
			fmt.Println("Thanks for using the app!")
			return
			//break
		}
	}	
}

func writeToafile(bal float64) {
	balVal := fmt.Sprint(bal)
	os.WriteFile(balanceFile, []byte(balVal), 0644)
}

func getBalFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}

	balText := string(data)
	balance, err := strconv.ParseFloat(balText, 64)

	if err !=nil {
		return 1000, errors.New("Failed to get balance value")
	}

	return balance, nil
}