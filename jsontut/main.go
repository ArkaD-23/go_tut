package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isadult"`
}

func main() {
	Person1 := Person{Name: "John", Age: 34, IsAdult: true}
	fmt.Println("The user data is :", Person1)

	//Marshalling -> json encoding
	jsonData, err := json.Marshal(Person1)
	if err != nil {
		fmt.Println("Error is :", err)
	}
	fmt.Println("Person1 data in json format :", string(jsonData))

	//Unmarshalling -> json decoding
	var decodedData Person
	errors := json.Unmarshal(jsonData, &decodedData)
	if errors != nil {
		fmt.Println("Error is :", errors)
	}
	fmt.Println("Decoded data is :", decodedData)
}
