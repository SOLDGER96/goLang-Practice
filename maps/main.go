package main

import "fmt"

func main() {
	names := map[string]string{}
	names["tany"]="gawade"
	names["kundu"]="dhage"
	names["addy"]="shinde"


	fmt.Print(names)
	fmt.Println(names["ayush"])

	// Appending a Map
	names["pranav"] = "jadhav"
	fmt.Println(names)

	// Deleting a value in a Map
	delete(names, "hacker")
	fmt.Println(names)
}
