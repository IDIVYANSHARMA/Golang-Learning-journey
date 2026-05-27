package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define a struct to represent the data we want to write to the JSON file
type User struct {
	Name string
	Age int
}


func main() {

	// Create a new user
	user := User{
		Name: "Aman",
		Age: 30,
	}

	// Create a new Json file
	file, _ := os.Create("data.json")

	// Write data to the JSON file
	json.NewEncoder(file).Encode(user)
    
	// Close the file
	file.Close()

	fmt.Println("Data written to JSON file successfully!")
}