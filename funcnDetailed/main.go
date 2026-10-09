package main

import "fmt"

func main() {
	numbers := []int{1,2,3,4}
	dn := doubleNumber(&numbers)
	fmt.Println(numbers)
	fmt.Println(dn)
}

func doubleNumber(dn *[]int) []int {
	dnum := []int{}
	for _,v := range *dn {
		dnum = append(dnum, v*2)
	}
	return dnum
}