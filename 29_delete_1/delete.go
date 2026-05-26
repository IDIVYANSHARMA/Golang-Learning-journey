package main 

import (
	"fmt"
	"os"
)

func main() {
	// Delete the file
	err := os.Remove("new.txt") // os.Remove deletes the specified file

	if err != nil {
		fmt.Println("error found:", err) // If there is an error, print it and exit the function
		return
	}

	fmt.Println("file deleted successfully!")
}