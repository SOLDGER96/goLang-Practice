package main
import (
	"fmt"
)

func main() {

	var accBal = 5000.0
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

	

	if choice == 1 {
		fmt.Println("Your Account Balance: ", accBal)

		}else if choice == 2 {
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

		}else if choice == 3 {
			fmt.Print("Enter Withdrawl Amount: ")
			var w float64
			fmt.Scan(&w)
			if w <= 0 {
				fmt.Print("Invalid Amount, Must be greater then 0")
				//return
				continue
			}
			if w > accBal {
				fmt.Print("Invalid Amount, Not enough Balance!")
				//return
				continue
			}
			accBal -= w
			fmt.Println("Your Balance: ", accBal)

		}else {
			fmt.Println("Thank you!")
			//return
			break
		}
	}

	fmt.Println("Thanks for using the app!")
}
