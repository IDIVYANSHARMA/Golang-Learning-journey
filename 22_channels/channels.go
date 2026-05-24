package main 

import "fmt"

func sendData(ch chan string) {
	ch <- "hello Programmer"
}

func main() {
	ch := make(chan string)
	go sendData(ch) 

	data := <-ch    // 
	fmt.Println(data)
}