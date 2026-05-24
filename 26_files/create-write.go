package main 

import (
	"fmt"
	"os"
)

func  main() {
	// Create a new file
	file, err := os.Create("hello.txt")  // os.Create creates the file if it doesn't exist, or truncates it if it does

	if err != nil {
		fmt.Println("Error found:", err) // If there is an error, print it and exit the function
		return
	}
	
	defer file.Close() // Ensure the file is closed when we're done

	// Write some text to the file
	file.WriteString("Hello Golang Expert :)")  // WriteString writes the specified string to the file
	fmt.Println(file.Name())
	fmt.Println("File created and text written successfully!")
}