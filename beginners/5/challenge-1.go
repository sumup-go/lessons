package main

import "fmt"

func countSheep(sheep []bool) int {
	count := 0
	for _, s := range sheep {
		if s {
			count++
		}
	}
	return count
}

var sheep = []bool{
	true, true, false, true,
	false, true, false, true,
	false, true, true, true,
	true, true, true, false,
}

func main() {
	fmt.Println(countSheep(sheep))
}
