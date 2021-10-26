package main

import (
	"fmt"
	"sync"
	"time"
)

func execute(command string, wg *sync.WaitGroup) string {
	defer wg.Done()
	// simulate some computation
	time.Sleep(time.Millisecond * 200)

	runes := []rune(command)
	for i := 0; i < len(command)/2; i++ {
		j := len(command) - 1 - i
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)
	var ans1, ans2 string
	go func() {
		ans1 = execute("do it", &wg)
	}()
	go func() {
		ans2 = execute("do it again", &wg)
	}()

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
	fmt.Println(ans1)
	fmt.Println(ans2)
}
