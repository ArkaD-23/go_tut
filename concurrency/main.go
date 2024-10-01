package main

import (
	"fmt"
	"time"
)

type Message struct {
	from    string
	payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) startAndListen() {
	// you can name your for loop
running:
	for {
		select {
		case msg := <-s.msgch:
			fmt.Printf("Message received from %s : %s\n", msg.from, msg.payload)
		case <-s.quitch:
			fmt.Println("The server is doing a graceful shutdown......")
			break running
		default:
		}
	}
	fmt.Println("The server is shut down......")
}

func (s *Server) sendMessage(payload string) {
	//fmt.Println("Sending message......")
	msg := Message{
		from:    "John Doe",
		payload: payload,
	}
	s.msgch <- msg
}

func serverShut(quitch chan struct{}) {
	close(quitch)
}

func main() {
	server := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go server.startAndListen()
	go func() {
		time.Sleep(2 * time.Second)
		server.sendMessage("Hello from Go!")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		serverShut(server.quitch)
	}()
}

// Synchronous and Asynchronous staff
// func main() {
// 	userId := 10
// 	now := time.Now()
// 	respch := make(chan string, 128)
// 	var wg sync.WaitGroup

// 	go fetchUser(userId, respch, &wg)
// 	go fetchUserRecommendations(userId, respch, &wg)
// 	go fetchUserLikes(userId, respch, &wg)
// 	wg.Add(3)
// 	wg.Wait()
// 	close(respch)

// 	// fmt.Println(user)
// 	// fmt.Println(userRecommendations)
// 	// fmt.Println(userLikes)

// 	for resp := range respch {
// 		fmt.Println(resp)
// 	}

// 	fmt.Println(time.Since(now)) //Output = 251ms without goroutines
// 	//Output = 121 ms after goroutines

// }

// func fetchUser(userId int, respch chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(80 * time.Millisecond)
// 	respch <- "user"
// }

// func fetchUserRecommendations(userId int, respch chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(120 * time.Millisecond)
// 	respch <- "user recommendations"
// }

// func fetchUserLikes(userId int, respch chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(50 * time.Millisecond)
// 	respch <- "user likes"
// }
