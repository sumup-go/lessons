package main

import (
	"fmt"
	"time"
)

func execute(command string) {
	// simulate some computation
	time.Sleep(time.Millisecond * 200)
}

func main() {
	start := time.Now()
	go execute("do it")
	go execute("do it again")
	fmt.Println(time.Now().Sub(start))
}
