package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	ch <- 1
	fmt.Println("one")
}


// Questions:
// What happens if no goroutine reads the received value?