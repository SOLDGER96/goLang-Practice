package main

import "fmt"

func main() {
	demoIntegers := []int{1,2,3,4}
	fmt.Println(demoIntegers)

	demoIntegers = append(demoIntegers, 5,6,7)
	fmt.Println(demoIntegers)

	// This can be used to delete the elemets in the slice
	demoIntegers = demoIntegers[3:]
	fmt.Println(demoIntegers)
}

// func main() {

// 	var productNames [4]string = [4]string{"Book", "Novel"}
// 	prices := [4]int{2, 4, 3, 7}

// 	// Slice ==> Index 1 is included and index 3 is excluded
// 	feaPrices := prices[1:3]
	
// 	fmt.Println(prices[3])
// 	fmt.Println(productNames)
// 	fmt.Println(feaPrices)


// 	sample := [5]int{1,2,3,4,5}
// 	sampleSlice := sample[2:]
// 	fmt.Println(sample)

// 	demoSlice := sampleSlice[:1]

// 	// Any modification in the slice updates the original array 
// 	// ==> It means slice does not create a copy, it works directly on the array
// 	sampleSlice[0] = 9
// 	fmt.Println(sample)

// 	// Use of len function and cap fuction
// 	// Here the len is 1 but cap 3 because demoSlice is build on top of another slice of size 3
// 	fmt.Println(len(demoSlice), cap(demoSlice))
// }
