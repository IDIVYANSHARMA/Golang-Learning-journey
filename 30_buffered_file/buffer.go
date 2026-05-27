package main 

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

  // Create a new file
	file, _ := os.Create("data.txt")

  // create a buffered writer
	writer := bufio.NewWriter(file)

  // Write data to the buffer
	writer.WriteString("Hello Golang Engineer! \n")
	writer.WriteString("this is a buffered file writing example.")

  // Flush the buffer to write data to the file
	writer.Flush()

  // Close the file
	file.Close()

	fmt.Println("Data written to file successfully!")
}