package main

import (
	"fmt"
)

func main() {
	fmt.Println(getNumber1())
}

func getNumber1() int {
	var i int
	done := false

	go func() {
		i = 5
		done = true
	}()

	for done != true {
	}
	return i
}

// Questions.
// Can we serialize access using flag `done`?
