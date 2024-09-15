package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cannot be zero.")
	}
	return a / b, nil
}

func main() {
	data, err := divide(10, 4)
	//error handling
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data is =", data)
	}
	// _ identifier is used if you dont want access any variable
	// Ex. data, _ := divide(10, 4)
}
