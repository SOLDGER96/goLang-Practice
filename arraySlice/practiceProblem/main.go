package main

import "fmt"

type Product struct {
	title	string
	id		string
	price	int64 	
}

func main() {
	hobby := [3]string{"run", "jump", "play"}
	
	fmt.Println(hobby)
	fmt.Println(hobby[0])
	hobby12 := hobby[1:]
	fmt.Println(hobby12)

	hobby01 := hobby[:2]
	fmt.Println(hobby01)


	fmt.Println("---------------------------")

	courseGoals := []string{"complete", "practice"}
	fmt.Println(courseGoals)

	courseGoals[1] = "inTime"
	fmt.Println(courseGoals)

	courseGoals = append(courseGoals, "code")
	fmt.Println(courseGoals)

	fmt.Println("---------------------------")

	products := []Product{
		Product{
			"title1",
			"abc123",
			756,
		},
		{
			"title2",
			"abc1234",
			758,
		},
	}
	fmt.Println(products)

	// Appending a structure
	newProduct := Product{
		"title3",
		"ab12",
		759,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line. 
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.