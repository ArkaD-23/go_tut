package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {

	// fmt.Println("First line.....")
	// fmt.Println("Middle line.....")
	// fmt.Println("Last line.....")
	//Output
	// First line.....
	// Middle line.....
	// Last line.....

	// fmt.Println("First line.....")
	// defer fmt.Println("Middle line.....")
	// fmt.Println("Last line.....")
	// Output
	// First line.....
	// Last line.....
	// Middle line.....

	fmt.Println("First line.....")
	data := add(5, 6)
	defer fmt.Println("data = ", data)
	defer fmt.Println("Middle 2 line.....")
	fmt.Println("Last line.....")
	// Output
	// First line.....
	// Last line.....
	// Middle 2 line.....
	// data =  11
	// when there are two or more defer then stack will be formed and that is at first will go in the statck first and execute at last
}
