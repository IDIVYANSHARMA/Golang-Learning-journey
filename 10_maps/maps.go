package main 

import "fmt"
import "maps"

func main() {
	m1 := map[string]int {"price": 100, "quantity": 10}

	m2 := map[string]int {"price": 100, "quantity": 14}
	
	fmt.Println(maps.Equal(m1, m2))
}