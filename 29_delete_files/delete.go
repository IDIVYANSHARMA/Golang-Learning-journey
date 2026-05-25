package main 

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// read the data from the file
	data, err := os.ReadFile("hello.txt")
	
	if err != nil {
		fmt.Println(err)
		return 
	}

	//convert bytes to string
	content := string(data)

	//Remove specific text from the content
	updatedContent := strings.Replace(
		content,
		"Welcome to the Go World!",
		"",
		1,
	)

	// Write the updated content back to the file
	err = os.WriteFile("hello.txt", []byte(updatedContent), 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Message removed successfully!")
}