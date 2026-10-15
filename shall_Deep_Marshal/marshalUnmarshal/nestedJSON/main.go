package main

import (
	"encoding/json"
	"fmt"
)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Address Address `json:"address"`
}

type Address struct{
	City string `json:"city"`
	Zipcode string `json:"zipcode"`
}



func main() {
	jsonBytes := []byte(`{"name": "Alice",
  	"age": 28,
  	"address": {
    	"city": "Pune",
    	"zipcode": "411001"}
		}`)
	
	var newUser User
	UnMarshal(jsonBytes, &newUser)
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
	fmt.Println(v.Address.City)
}