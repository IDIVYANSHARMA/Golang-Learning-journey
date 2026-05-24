package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(200 * time.Millisecond)
	}
}

func printLetters() {
	for ch := 'A'; ch <= 'D'; ch++ {
		fmt.Println(string(ch))
		time.Sleep(200 * time.Millisecond)  //
	}
}

func main() {
	go printNumbers() // Start printNumbers in a new goroutine
	go printLetters() // Start printLetters in a new goroutine

	time.Sleep(2 * time.Second)  // Wait for goroutines to finish before exiting the main function So we pause the main function for 2 seconds to allow the goroutines to complete their execution before the program exits.
}