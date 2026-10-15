# JSON Marshaling and Unmarshaling in Go

When you want to send Go structs over the internet or save them to a file, you have to convert them into a universal format like JSON. Go handles this using the built-in `encoding/json` package.

## 1. Struct Setup: The Privacy Rule and JSON Tags

Before you can convert a struct to JSON, you must follow the **Privacy Rule**: Go's JSON package can only "see" fields if they start with an uppercase letter (making them public). 

However, standard JSON keys are usually lowercase. To fix this, we attach **JSON tags** using backticks (`` ` ``) to tell Go exactly what to name the fields in the final JSON text.

```go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

## 2. Marshaling (Go Struct -> JSON)

**Marshaling** is the process of translating a Go struct into JSON text. The `json.Marshal()` function takes your struct and returns two things: the raw JSON data (as a byte slice) and an error.

Because it returns a raw `[]byte`, you must wrap it in `string()` to read it as standard text.

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	u := User{
		Name: "Tanmay",
		Age:  23,
	}

	jsonData, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	
	// Convert the raw byte slice to a readable string
	fmt.Println(string(jsonData)) 
	// Outputs: {"name":"Tanmay","age":23}
}
```

## 3. Unmarshaling (JSON -> Go Struct)

**Unmarshaling** is the exact opposite: taking JSON text from the internet and rebuilding it back into a Go struct. 

The `json.Unmarshal()` function requires two things:
1. The raw JSON byte slice.
2. A **pointer** (`&`) to an empty struct. The pointer is critical because it gives the function the exact memory address so it can directly modify and fill your struct in memory.

```go
func main() {
	jsonBytes := []byte(`{"name": "Alice", "age": 28}`)
	var newUser User
	
	// Pass the pointer (&newUser) so Unmarshal can modify it directly
	err := json.Unmarshal(jsonBytes, &newUser)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}
	
	fmt.Println(newUser) 
	// Outputs: {Alice 28}
}
```

## 4. Deep / Nested JSON

Real-world APIs often return **nested JSON** (objects inside of objects). To parse this, your Go structs must perfectly mirror the nested shape. 

Be careful to match data types exactly! For example, if a number in the JSON has quotes around it (`"zipcode": "411001"`), Go expects a `string`, not an `int`.

```go
// 1. Create a struct for the nested object
type Address struct {
	City    string `json:"city"`
	Zipcode string `json:"zipcode"` // string because JSON has quotes!
}

// 2. Include the new struct inside the parent struct
type User struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	jsonBytes := []byte(`{
		"name": "Alice", 
		"age": 28, 
		"address": {
			"city": "Pune", 
			"zipcode": "411001"
		}
	}`)
	
	var newUser User
	json.Unmarshal(jsonBytes, &newUser)
	
	// Use "dot notation" to drill down into the nested structs
	fmt.Println(newUser.Address.City) 
	// Outputs: Pune
}
```