package main

import (
	"fmt"
)

var p1, p2 string

func main() {
	fmt.Scanf("%s\n", &p1)
	fmt.Scanf("%s\n", &p2)
	if p1 == p2 {
		fmt.Println("Draw!")
	}
	switch p1 {
	case "paper":
		if p2 == "scissors" {
			fmt.Println("Player 2 won!")
		} else {
			if p2 == "rock" {
				fmt.Println("Player 1 win")
			}
		}
	case "rock":
		if p2 == "scissors" {
			fmt.Println("Player 1 won!")
		}

	case "scissors":
		if p2 == "rock" {
			fmt.Println("Player 2 won!")
		}
	}
}
