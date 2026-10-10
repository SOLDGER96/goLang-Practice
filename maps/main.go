package main

import "fmt"

func main() {
	names := map[string]string{}
	names["tany"] = "gawade"
	names["kundu"] = "dhage"
	names["addy"] = "shinde"

	fmt.Print(names)
	fmt.Println(names["ayush"])

	// Appending a Map
	names["pranav"] = "jadhav"
	fmt.Println(names)

	// Deleting a value in a Map
	delete(names, "hacker")
	fmt.Println(names)
	// named := map[string]int{
	// 	"tgawade": 0,
	// 	"adit":    1,
	// }
	// if v, ok := named["tgawad"]; ok {
	// 	fmt.Println("value exists", v)
	// }else {
	// 	fmt.Println("value does not exist")
	// }

	// if _,err := tanny <=addy ; err!=nil{

	// }
	// p := 10  <=11
	// if p {
	// 	fmt.Print(v)
	// }
}
