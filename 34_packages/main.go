package main 

import (
	"fmt"
	"myproject/calculator"
)

func main() {

	sum := calculator.Add(5, 15)
	sub := calculator.Subtract(15, 5)

	fmt.Println("Addition: ", sum)
	fmt.Println("Subtraction: ", sub)
}