package main

import (
	"fmt"
	"time"
)

func execute(command string, out chan string) {
	time.Sleep(time.Millisecond * 200)

	runes := []rune(command)
	for i := 0; i < len(command)/2; i++ {
		j := len(command) - 1 - i
		runes[i], runes[j] = runes[j], runes[i]
	}

	out <- string(runes)
}

func main() {
	start := time.Now()
	commands := []string{
		"do it",
		"do it please",
		"do some more",
	}

	out := make(chan string)
	for _, command := range commands {
		go execute(command, out)
	}

	for i := 0; i < len(commands); i++ {
		fmt.Println(<-out)
	}

	fmt.Println(time.Now().Sub(start))
}
