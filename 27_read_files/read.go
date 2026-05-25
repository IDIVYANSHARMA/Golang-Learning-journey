package main 

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("hello.txt") // os.ReadFile reads the entire content of the specified file and returns it as a byte slice

	if err != nil {
		fmt.Println("Error found:", err) // If there is an error, print it and exit the function
		return
	}
	fmt.Println("File contents:", string(data)) // Convert the byte slice to a string and print it
}