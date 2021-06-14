package main

import "fmt"

type (
	Cat   string
	Dog   int
	Snake struct{}
)

func (c Cat) Speak() string {
	return "Meow"
}

func (d Dog) Speak() string {
	return "Woof"
}

func (s Snake) Speak() string {
	return "Hiss"
}

type Speaker interface {
	Speak() string
}

func speak(sp Speaker) {
	fmt.Println(sp.Speak())
}

func main() {
	speak(Cat("")) // Meow
	speak(Dog(0))  // Woof
	speak(Snake{}) // Hiss
}
