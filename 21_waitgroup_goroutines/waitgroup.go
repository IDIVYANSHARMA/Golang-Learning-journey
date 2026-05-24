package main 

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it finishes
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(200 * time.Millisecond)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it finishes
	for ch := 'A'; ch <= 'D'; ch++ {
		fmt.Println(string(ch))
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup

	wg.Add(2) // We have 2 goroutines to wait for

	go printNumbers(&wg) // Start printNumbers in a new goroutine
	go printLetters(&wg) // Start printLetters in a new goroutine

	wg.Wait() // Wait for both goroutines to finish before exiting the main function
}