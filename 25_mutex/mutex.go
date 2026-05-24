package main 

import (
	"fmt"
	"sync"
)

var count = 0
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done() // signal that this goroutine is done

	mutex.Lock()     //lock the shared resource
	count++
	mutex.Unlock()   //unlock the shared resource
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1) // increment the WaitGroup counter
		go increment(&wg) // start a new goroutine to increment the count
	}
	wg.Wait() // wait for all goroutines to finish
	fmt.Println("Final count:", count) // print the final count
}