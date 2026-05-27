package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	// Create a new CSV file
	file, _ := os.Create("data.csv")

	// Create a new CSV writer
	writer := csv.NewWriter(file)

	// Write data to the CSV file
	data := [][]string {
		{"Name", "Age", "City"},
		{"aman", "30", "New York"},
		{"sara", "25", "Los Angeles"},
	}
    
	// Write all data to the CSV file
	writer.WriteAll(data)

	// Flush the writer to write data to the file
	writer.Flush()

	// Close the file
    file.Close()

	fmt.Println("Data written to CSV file successfully!")
}