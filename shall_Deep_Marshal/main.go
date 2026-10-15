package main

import "fmt"

func main() {
	sliceA := []string{"a", "b"}
	sliceB := sliceA
	// sliceA = append(sliceA, "c")
	sliceB[0] = "z"
	fmt.Println(sliceA)
	fmt.Println(sliceB)

	sliceC := make([]string, len(sliceA))
	fmt.Println(sliceC)
	copy(sliceC, sliceA)
	fmt.Println(sliceC)
}