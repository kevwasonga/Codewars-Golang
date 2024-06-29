package main

import "fmt"

func main() {
	fmt.Println(Rps("scissors", "paper")) // Player 1 won!
	fmt.Println(Rps("paper", "scissors")) // Player 2 won!
	fmt.Println(Rps("paper", "paper"))    // Draw!
}

func Rps(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}

	switch p1 {
	case "rock":
		if p2 == "paper" {
			return "Player 2 won!"
		}
		return "Player 1 won!"
	case "paper":
		if p2 == "scissors" {
			return "Player 2 won!"
		}
		return "Player 1 won!"

	case "scissors":
		if p2 == "paper" {
			return "Player 1 won!"
		}
		return "Player 2 won!"
	}
	return "invalid"
}
