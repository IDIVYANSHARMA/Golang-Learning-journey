package main 

import (
	"fmt"
	"os"
)

func main() {
	// Open the file for reading
	file, err := os.OpenFile(
		"hello.txt",
		os.O_APPEND | os.O_WRONLY, // Open the file in append and write-only mode
		0644, // Set the file permissions (rw-r--r--)
	)

	if err != nil {
		fmt.Println("error found:", err) // If there is an error, print it and exit the function
		return
	}

	defer file.Close() // Ensure the file is closed when we're done

	// Append some text to the file
	file.WriteString("\n Welcome to the Go World!") // WriteString appends the specified string to the file
	fmt.Println("Text appended successfully!")
}