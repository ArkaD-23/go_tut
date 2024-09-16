package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating file.....")
	// 	return
	// }
	// defer file.Close()
	// fmt.Println("File created successfully.....")

	// text := "This is just test"
	// _, errors := io.WriteString(file, text) //returns butes written
	// if errors != nil {
	// 	fmt.Println("Error while writing in file.....")
	// 	return
	// }
	// fmt.Println("Writing successful.....")

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error while opening file.....")
		return
	}
	defer file.Close()

	//creating a buffer
	buffer := make([]byte, 1024)

	//read the file into buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file.....", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}
}
