package main

import "fmt"

type strMap map[string]string

func (s strMap) output() {
	fmt.Println(s)
}

func main() {
	names := make([]string, 2, 5)
	
	names[0] = "a1"
	names[1] = "a2"
	names = append(names, "a3")

	fmt.Println(names)

	fmt.Println("-------------------")

	// make for maps
	nameSurname := make(strMap, 5)

	nameSurname["tanmay"] = "gawade"
	nameSurname["ayush"] = "parab"
	nameSurname["vasu"] = "chile"
	
	nameSurname.output()

	// for loop for array, slice, map
	for index, value := range nameSurname {
		fmt.Printf("Name: %v, Surname: %v\n",index, value)
	}
}
