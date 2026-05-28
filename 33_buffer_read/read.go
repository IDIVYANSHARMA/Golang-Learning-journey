package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file for reading
	file, err := os.Open("data.txt")

	// Check for errors while opening the file
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line and print each line
	for scanner.Scan() {
		Line := scanner.Text()
		fmt.Println(Line)
	}

	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Close the file
	file.Close()
}
