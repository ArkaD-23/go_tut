package main

import (
	"fmt"
	"time"
)

type Server struct {
	users map[string]string
	//userch chan string
	quitch chan struct{}
}

func NewServer() *Server {
	return &Server{
		users: make(map[string]string),
		//userch: make(chan string),
		quitch: make(chan struct{}),
	}
}

// func (s *Server) Start() {
// 	go s.loop()
// }

//func (s *Server) loop() {
//	for {
// user := <-s.userch
// s.users[user] = user
// fmt.Printf("Adding new user %s\n", user)
// 		select {
// 		case msg := <-s.userch:
// 			fmt.Println(msg)
// 		case <-s.quitch:
// 			return
// 		}
// 	}
// }

func main() {
	userch := make(chan string, 1)

	userch <- "BOB"
	//userch <- "ALICE"
	// //userch <- "JOHN" // waiting for the consume of the channel

	// for unbuffered channels
	// go func() {
	// 	userch <- "BOB"
	// }()

	// user := <-userch
	// fmt.Println(user)
	// senddMessage(userch)
	// readMessage(userch)

	go func() {
		time.Sleep(2 * time.Second)
		close(userch)
	}()

	username, ok := <-userch
	if !ok {
		fmt.Println("user chan closed returning")
		return
	}
	fmt.Printf("Username: %s\n", username)

	if len(username) == 0 {
		panic("User not good")
	}
}

// Protecting channels

// func senddMessage(msgch chan<- string) {
// 	msgch <- "Hello!"
// }

// func readMessage(msgch <-chan string) {
// 	msg := <-msgch
// 	fmt.Println(msg)
// }
