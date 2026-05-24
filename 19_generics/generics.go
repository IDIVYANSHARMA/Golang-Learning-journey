package main

import "fmt"

func add[T int | float64](a T, b T) T { // “T can be int OR float64”  T is called a type parameter.
	return a + b
}

func main() {
	fmt.Println(add(5, 15))
	fmt.Println(add(3.5, 4.5))
}