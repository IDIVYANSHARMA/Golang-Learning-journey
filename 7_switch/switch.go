package main 

import "fmt"
import "time"

func main() {
	// i := 3
	// switch i {
	// case 1 :
	// 	fmt.Println("one")
	// case 2 :
	// 	fmt.Println("two")
	// case 3 :
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("others")
	// }
	switch time.Now().Weekday() {
    case time.Saturday, time.Sunday :
		fmt.Println("its Weekend")
	default:
		fmt.Println("its workday")
	}
}