package main

import "fmt"

func main() {
	userch := make(chan string, 1)

	// userch <- "BOB"
	// userch <- "ALICE"
	// //userch <- "JOHN" // waiting for the consume of the channel

	// for unbuffered channels
	// go func() {
	// 	userch <- "BOB"
	// }()

	// user := <-userch
	// fmt.Println(user)
	senddMessage(userch)
	readMessage(userch)
}

// Protecting channels

func senddMessage(msgch chan<- string) {
	msgch <- "Hello!"
}

func readMessage(msgch <-chan string) {
	msg := <-msgch
	fmt.Println(msg)
}
