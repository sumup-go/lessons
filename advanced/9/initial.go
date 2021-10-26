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
	execute("do it")
	execute("do it again")
	fmt.Println(time.Now().Sub(start))
}
