package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Requesting to web.....")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while connecting.....")
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response:", string(data))

}
