package main

import "fmt"

func psr(p1 string, p2 string) string {
	if p1 == p2 {
		return "Tie!"
	}

	if p1 == "paper" {
		if p2 == "rock" {
			return "Player 1 wins!"
		}
		if p2 == "scissors" {
			return "Player 2 wins!"
		}
	} else if p1 == "scissors" {
		if p2 == "rock" {
			return "Player 2 wins!"
		}
		if p2 == "paper" {
			return "Player 1 wins!"
		}
	} else if p1 == "rock" {
		if p2 == "paper" {
			return "Player 2 wins!"
		}
		if p2 == "scissors" {
			return "Player 1 wins!"
		}
	}

	return "Impossible game!"
}

func main() {
	fmt.Println(psr("paper", "paper"))
	fmt.Println(psr("paper", "scissors"))
	fmt.Println(psr("rock", "scissors"))
	fmt.Println(psr("paper", "rock"))
}
