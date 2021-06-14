package main

import "fmt"

var hello = "something"

func main() {
	a := 1
	{
		b := 2
		fmt.Println(a, b)
	}
	// fmt.Println(a, b)
}

func slices() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice)

	mySlice := []int{1, 2, 3, 4}
	mySlice = append(mySlice, 5, 6, 7)

	type friend struct {
		Name    string
		Country string
		Pets    int
		IsClose bool
	}

	_ = []friend{}
}
