package main

import (
	"encoding/json"
	"fmt"
)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}



func main() {
	u := User{
		Name: "tanmay",
		Age: 23,
	}

	marshal(u)

	//For UnMarshal

	var newuser User
	jsonBytes := []byte(`{"name": "Alice", "age": 28}`)
	UnMarshal(jsonBytes, &newuser)
	
	fmt.Println(newuser)
}

func marshal(data User) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error!")
		return
	}
	
	fmt.Println(string(jsonData))
}

func UnMarshal(data []byte, v *User) {
	err := json.Unmarshal(data, v)
	if err != nil {
		return
	}
	fmt.Println(*v)
}