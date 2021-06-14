package main

import "fmt"

func rev(str string) string {
	var reversed string

	idx := len(str) - 1
	for idx >= 0 {
		reversed += string(str[idx])
		idx--
	}

	return reversed
}

func main() {
	fmt.Println(rev("use the force")) // "ecrof eht esu"
	fmt.Println(rev("hello world"))   // "dlrow olleh"
}
