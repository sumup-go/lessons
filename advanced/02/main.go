package main

import "fmt"

func main() {
	fmt.Println(double(2))
}

func double(i int) int {
	return i * 2
}
