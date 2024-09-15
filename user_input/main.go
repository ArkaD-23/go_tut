package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name.....")
	//var name string
	// fmt.Scan(&name)
	// fmt.Printf("Hello, Mr.%s\n", name)
	// Scan takes the value till the immediate white space so we use bufio package
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, Mr.%s\n", name)
}
