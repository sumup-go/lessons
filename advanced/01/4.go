package main

import (
	"fmt"
)

func f1(out chan int) {
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			out <- i
		}
	}()
}

func f2(inp chan int) {
	for v := range inp {
		fmt.Println(v)
	}
}

func pipeline() {
	ch := make(chan int)
	f1(ch)
	f2(ch)
}

func main() {
	pipeline()
}
