package main

import "fmt"

func main() {
	fmt.Println(fact(5))
}

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num - 1)
	// result := 1
	// for i := 1; i <= num; i++ {
	// 	result = result * i
	// }
	// return result
}