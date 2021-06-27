package main

import (
	"fmt"
)

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()

	fmt.Println("a")

	return i
}

// Questions:
// 1. The output of the program.
// 2. What happens if we run the program with a race detector?
// 2. How can we serialize access to variable `i`?
