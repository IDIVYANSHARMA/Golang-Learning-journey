package main 

import "fmt"

type orderStatus int

const (
	Recieved orderStatus = iota  // iota is a special Go keyword that automatically increments the value of the constant, starting from 0
	Processing
	Shipped 
	Delivered
)

func main() {
	var status orderStatus = Shipped
	fmt.Println(status)
}