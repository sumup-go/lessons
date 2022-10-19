package main

import (
	"fmt"
)

func do(ch chan bool) {
	<-ch
	fmt.Println("got signal")
}

func main() {
	ch := make(chan bool)
	go do(ch)

	ch <- true
}
