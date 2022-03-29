package main

import (
	"fmt"
	"sync"
	"time"
)

func execute(command string, wg *sync.WaitGroup) {
	defer wg.Done()
	// simulate some computation
	time.Sleep(time.Millisecond * 200)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)
	go execute("do it", &wg)
	go execute("do it again", &wg)

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
