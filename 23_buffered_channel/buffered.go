package main 

import "fmt"

func sendData(ch chan string) {
	ch <- "hello Programmer" 
}

func main() {
	ch := make(chan string, 2) // buffered channel with capacity of 2
	go sendData(ch)
	ch <- "welcome to Go programming" // sending another message to the buffered channel

	data1 := <-ch  // receiving the first message
	data2 := <-ch   // receiving the second message

	fmt.Println(data1)
	fmt.Println(data2)

	// ch := make(chan int, 2)

	// ch <- 10
	// ch <- 20

	// data1 := <-ch  // receiving the first message
	// data2 := <-ch   // receiving the second message

	// fmt.Println(data1)
	// fmt.Println(data2)
}