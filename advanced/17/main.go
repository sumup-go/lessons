package main

import "fmt"

func main() {
	fmt.Println(Reverse("Hello World"))
}

func Reverse(s string) string {
	var out []rune
	for _, b := range s {
		out = append([]rune{b}, out...)
	}

	return string(out)
}
