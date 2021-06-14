package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) string {
	names := strings.Split(name, " ")
	return fmt.Sprintf("%c.%c", names[0][0], names[1][0])
}

func main() {
	fmt.Println(getInitials("John Smith"))
	fmt.Println(getInitials("Bob Ross"))
}
