package main 

import "fmt" 

func sendNumbers(numch chan int) {
	numch <- 10
	numch <- 20
}

func sendMessages(msgch chan string) {
	msgch <- "hello"
	msgch <- "Golang Expert"
}

func main() {
	numch := make(chan int, 2) // buffered channel for integers
	msgch := make(chan string, 2) // buffered channel for strings

	go sendNumbers(numch) // sending numbers in a separate goroutine
	go sendMessages(msgch) // sending messages in a separate goroutine

	num1 := <- numch // receiving the first number
	num2 := <- numch // receiving the second number

	msg1 := <- msgch // receiving the first message
	msg2 := <- msgch // receiving the second message

	fmt.Println("Recieved numbers:", num1, num2)
	fmt.Println("Recieved messages:", msg1, msg2)
}