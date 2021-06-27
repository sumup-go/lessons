package main

import (
	"fmt"
)

func main() {
	fmt.Println(getNumber2())
}

func getNumber2() int {
	var i int
	go func() {
		i = 5
	}()
	
	fmt.Println("a")
	
	return i
}

// Questions:
// 1. Will the output change if we add a print statement?
