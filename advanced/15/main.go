package main

import (
	"fmt"
)

type Stinger interface {
	String() string
}

type person struct {
	name string
	age  int
}

func (p *person) changeName() {
	p.name = "Frederik"
}

func (p *person) String() string {
	return p.name
}

type changer interface {
	changeName()
	String() string
}

type mycomp interface {
	int
	string
}

type any = interface{}

func main() {
	// p := mycomp(&person{"Sam", 32})
	// p.changeName()
	// fmt.Println(p)

	var a interface{} = nil
	fmt.Println(a == nil)

	var b interface{} = new(int)
	fmt.Println(b == nil)

	whatWillHappen()
}

func whatWillHappen() {
	// (type, data)

	var (
		a interface{} = []int{1}
		b interface{} = []int{2}

		// a interface{} = "string1"
		// b interface{} = "string1"
	)

	fmt.Println(b == a)
}
