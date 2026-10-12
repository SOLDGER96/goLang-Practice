package main

import "fmt"

// Custom Type
type transformFn func(int) int

func main() {
	numbers := []int{1,2,3,4}
	dn := doubleNumber(&numbers)
	fmt.Println("Original List: ", numbers)
	fmt.Println("Doubled List: ", dn)

	doubled := transformNum(&numbers, double)
	fmt.Println("Transformed into Double: ", doubled)

	tripled := transformNum(&numbers, triple)
	fmt.Println("Transformed into Tripple: ", tripled)
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

func doubleNumber(dn *[]int) []int {
	dnum := []int{}
	for _,v := range *dn {
		dnum = append(dnum, double(v))
	}
	return dnum
}

func transformNum(num *[]int, transform transformFn) []int {
	dnum := []int{}
	for _,v := range *num {
		dnum = append(dnum, transform(v))
	}
	return dnum
}