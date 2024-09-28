package main

import "fmt"

type Message struct {
	from    string
	payload string
}

type Server struct {
	msgch chan Message
}

func (s *Server) startAndListen() {
	for {
		msg := <-s.msgch
		fmt.Printf("Message received from %s : %s\n", msg.from, msg.payload)
	}
}

func (s *Server) sendMessage(payload string) {
	fmt.Println("Sending message......")
	msg := Message{
		from:    "John Doe",
		payload: payload,
	}
	s.msgch <- msg
}

func main() {
	server := &Server{
		msgch: make(chan Message),
	}
	go server.startAndListen()
	server.sendMessage("Hello from Go!")
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
