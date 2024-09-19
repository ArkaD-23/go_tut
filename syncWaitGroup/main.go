package main

import (
	"fmt"
	"sync"
)

func worker(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started.....\n", n)
	fmt.Printf("Worker %d ended.....\n", n)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("Worker functions started.....")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Worker functions completed.....")
}
