package main

import "fmt"

func main() {
	num := []int{1,2,3,4,5}

	doubleNum := transformNum(&num, doubele)
	fmt.Println(doubleNum)

	trippleNum := transformNum(&num, func(num int) int {return num*3})
	fmt.Println(trippleNum)
}

func transformNum(num *[]int, transform func(int) int) []int {
	dnum := []int{}
	for _, val := range *num {
		dnum = append(dnum, transform(val))
	}
	return dnum
}

func doubele(num int) int {
	return num * 2
}