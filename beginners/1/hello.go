package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var stopWords = []string{"stop", "quit", "q"}

func main() {
	for {
		age := prompt("What is your age?")
		for _, w := range stopWords {
			if age == w {
				os.Exit(0)
			}
		}

		reply := PrintName(age)
		fmt.Println(reply)
	}
}

func prompt(phrase string) string {
	fmt.Println(phrase)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func PrintName(age string) string {
	if s, err := strconv.Atoi(age); err != nil {
		return "error"
	} else if s < 50 {
		return "You're too young!"
	} else {
		name := prompt("What is your name?")
		return name + " you are too old."
	}
}
